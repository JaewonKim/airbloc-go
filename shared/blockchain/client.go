package blockchain

import (
	"context"
	"github.com/airbloc/airbloc-go/shared/blockchain/bind"
	"github.com/airbloc/airbloc-go/shared/key"
	"github.com/airbloc/logger"
	ethbind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type Client struct {
	*ethclient.Client
	ctx        context.Context
	cfg        ClientOpt
	transactor *bind.TransactOpts
	contracts  *ContractManager
	logger     *logger.Logger
}

func NewClient(key *key.Key, rawurl string, cfg ClientOpt) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log := logger.New("ethereum")

	log.Debug(rawurl)
	// URL validation
	l, err := url.Parse(rawurl)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid URL: %s", rawurl)
	}
	if l.Scheme != "ws" {
		log.Error("Warning: You're using {} endpoint for Ethereum. Using WebSocket is recommended.",
			strings.ToUpper(l.Scheme))
	}

	// try to connect to Ethereum
	ethClient, err := ethclient.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	cid, err := ethClient.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	log.Info("Using {} network", getChainName(cid))

	client := &Client{
		Client: ethClient,
		ctx:    context.TODO(),
		cfg:    cfg,
		logger: log,
	}

	cm := NewContractManager(client)
	if err := cm.Load(cfg.DeploymentPath); err != nil {
		return nil, err
	}
	client.contracts = cm
	client.SetAccount(key)
	return client, nil
}

func (c Client) Account() *bind.TransactOpts {
	return c.transactor
}

func (c *Client) SetAccount(key *key.Key) {
	c.transactor = bind.NewKeyedTransactor(key.PrivateKey)
}

func (c *Client) GetContract(contractType interface{}) interface{} {
	contract := c.contracts.GetContract(contractType)
	if contract == nil {
		panic("Contract not registered: " + reflect.ValueOf(contractType).Type().Name())
	}
	return contract
}

func (c *Client) waitConfirmation(ctx context.Context) error {
	ch := make(chan *types.Header)
	sub, err := c.SubscribeNewHead(c.ctx, ch)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for count := c.cfg.Confirmation; count > 0; {
		select {
		case <-ch:
			count--
		case <-ctx.Done():
			return context.DeadlineExceeded
		}
	}
	return err
}

// Wait Mined
func (c *Client) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	methodName, details := GetTransactionDetails(c.contracts, tx)
	timer := c.logger.Timer()

	receipt, err := ethbind.WaitMined(ctx, c, tx)
	if err != nil {
		return nil, err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		timer.End("Transaction to {} failed", methodName, details)
		return nil, ErrTxFailed
	}
	timer.End("Transacted {}", methodName, details)
	// err = c.waitConfirmation(ctx)
	return receipt, err
}

// Wait Deployed
func (c *Client) WaitDeployed(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	if tx.To() != nil {
		return nil, ErrTxNoContract
	}

	receipt, err := c.WaitMined(ctx, tx)
	if err != nil {
		return nil, err
	}
	if receipt.ContractAddress == (common.Address{}) {
		return nil, ErrZeroAddress
	}

	code, err := c.CodeAt(ctx, receipt.ContractAddress, nil)
	if err == nil && len(code) == 0 {
		err = bind.ErrNoCodeAfterDeploy
	}
	// err = c.waitConfirmation(ctx)
	return receipt, err
}

package usecase

import (
	"database/sql"
	"testeAula/internal/order/entity"
	database "testeAula/internal/order/infra/dataBase"
	"testing"

	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type CalculatePriceUseCaseTestSuite struct {
	suite.Suite
	OrderRepository database.OrderRepository
	Db              *sql.DB
}

func (suite *CalculatePriceUseCaseTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	_, err = db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
	suite.OrderRepository = *database.NewOrderRepository(db)
}

func (suite *CalculatePriceUseCaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CalculatePriceUseCaseTestSuite))
}

func (suite *CalculatePriceUseCaseTestSuite) TestGivenAnOrder_WhenCalculatePrice_ThenIShouldBeCalculated() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	order.CalculatePrice()
	
	CalculatePriceInput:= OrderInputDTO{
		ID: 	order.ID,
		Price: 	order.Price,
		Tax: 	order.Tax,
	}

	CalculatePriceUseCase := NewCalculatePriceUseCase(suite.OrderRepository)
	output, err := CalculatePriceUseCase.Execute(CalculatePriceInput)
	suite.NoError(err)

	suite.Equal(order.ID, output.ID)
	suite.Equal(order.Price, output.Price)
	suite.Equal(order.Tax, output.Tax)
	suite.Equal(order.FinalPrice, output.FinalPrice)
}

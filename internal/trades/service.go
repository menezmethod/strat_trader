// Code generated by sqlc-grpc (https://github.com/walterwanderley/sqlc-grpc). DO NOT EDIT.

package trades

import (
	"context"
	"database/sql"
	"fmt"

	"go.uber.org/zap"

	pb "strat_trader/api/trades/v1"
	"strat_trader/internal/validation"
)

type Service struct {
	pb.UnimplementedTradesServiceServer
	logger  *zap.Logger
	querier *Queries
}

func (s *Service) CreateCountry(ctx context.Context, req *pb.CreateCountryRequest) (*pb.CreateCountryResponse, error) {
	name := req.GetName()

	result, err := s.querier.CreateCountry(ctx, name)
	if err != nil {
		s.logger.Error("CreateCountry sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.CreateCountryResponse{Country: toCountry(result)}, nil
}

func (s *Service) CreateCurrency(ctx context.Context, req *pb.CreateCurrencyRequest) (*pb.CreateCurrencyResponse, error) {
	var arg CreateCurrencyParams
	arg.Code = req.GetCode()
	arg.Name = req.GetName()
	arg.IsActive = req.GetIsActive()
	arg.IsBaseCurrency = req.GetIsBaseCurrency()

	result, err := s.querier.CreateCurrency(ctx, arg)
	if err != nil {
		s.logger.Error("CreateCurrency sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.CreateCurrencyResponse{Currency: toCurrency(result)}, nil
}

func (s *Service) CreateCurrencyRate(ctx context.Context, req *pb.CreateCurrencyRateRequest) (*pb.CreateCurrencyRateResponse, error) {
	var arg CreateCurrencyRateParams
	arg.CurrencyID = req.GetCurrencyId()
	arg.BaseCurrencyID = req.GetBaseCurrencyId()
	arg.Rate = req.GetRate()
	if v := req.GetTs(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid Ts: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
		arg.Ts = v.AsTime()
	} else {
		err := fmt.Errorf("field Ts is required%w", validation.ErrUserInput)
		return nil, err
	}

	result, err := s.querier.CreateCurrencyRate(ctx, arg)
	if err != nil {
		s.logger.Error("CreateCurrencyRate sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.CreateCurrencyRateResponse{CurrencyRate: toCurrencyRate(result)}, nil
}

func (s *Service) CreateCurrencyUsed(ctx context.Context, req *pb.CreateCurrencyUsedRequest) (*pb.CreateCurrencyUsedResponse, error) {
	var arg CreateCurrencyUsedParams
	arg.CountryID = req.GetCountryId()
	arg.CurrencyID = req.GetCurrencyId()
	if v := req.GetDateFrom(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid DateFrom: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
		arg.DateFrom = v.AsTime()
	} else {
		err := fmt.Errorf("field DateFrom is required%w", validation.ErrUserInput)
		return nil, err
	}
	if v := req.GetDateTo(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid DateTo: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
		if t := v.AsTime(); !t.IsZero() {
			arg.DateTo.Valid = true
			arg.DateTo.Time = t
		}
	}

	result, err := s.querier.CreateCurrencyUsed(ctx, arg)
	if err != nil {
		s.logger.Error("CreateCurrencyUsed sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.CreateCurrencyUsedResponse{CurrencyUsed: toCurrencyUsed(result)}, nil
}

func (s *Service) CreateCurrentInventory(ctx context.Context, req *pb.CreateCurrentInventoryRequest) (*pb.CreateCurrentInventoryResponse, error) {
	var arg CreateCurrentInventoryParams
	arg.TraderID = req.GetTraderId()
	arg.ItemID = req.GetItemId()
	arg.Quantity = req.GetQuantity()

	result, err := s.querier.CreateCurrentInventory(ctx, arg)
	if err != nil {
		s.logger.Error("CreateCurrentInventory sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.CreateCurrentInventoryResponse{CurrentInventory: toCurrentInventory(result)}, nil
}

func (s *Service) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	var arg CreateItemParams
	arg.Code = req.GetCode()
	arg.Name = req.GetName()
	arg.IsActive = req.GetIsActive()
	arg.CurrencyID = req.GetCurrencyId()
	if v := req.GetDetails(); v != nil {
		arg.Details = sql.NullString{Valid: true, String: v.Value}
	}

	result, err := s.querier.CreateItem(ctx, arg)
	if err != nil {
		s.logger.Error("CreateItem sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.CreateItemResponse{Item: toItem(result)}, nil
}

func (s *Service) CreateOffer(ctx context.Context, req *pb.CreateOfferRequest) (*pb.CreateOfferResponse, error) {
	var arg CreateOfferParams
	arg.TraderID = req.GetTraderId()
	arg.ItemID = req.GetItemId()
	arg.Quantity = req.GetQuantity()
	arg.Buy = req.GetBuy()
	arg.Sell = req.GetSell()
	if v := req.GetPrice(); v != nil {
		arg.Price = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetTs(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid Ts: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
		arg.Ts = v.AsTime()
	} else {
		err := fmt.Errorf("field Ts is required%w", validation.ErrUserInput)
		return nil, err
	}
	arg.IsActive = req.GetIsActive()

	result, err := s.querier.CreateOffer(ctx, arg)
	if err != nil {
		s.logger.Error("CreateOffer sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.CreateOfferResponse{Offer: toOffer(result)}, nil
}

func (s *Service) CreatePrice(ctx context.Context, req *pb.CreatePriceRequest) (*pb.CreatePriceResponse, error) {
	var arg CreatePriceParams
	arg.ItemID = req.GetItemId()
	arg.CurrencyID = req.GetCurrencyId()
	arg.Buy = req.GetBuy()
	arg.Sell = req.GetSell()
	if v := req.GetTs(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid Ts: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
		arg.Ts = v.AsTime()
	} else {
		err := fmt.Errorf("field Ts is required%w", validation.ErrUserInput)
		return nil, err
	}

	result, err := s.querier.CreatePrice(ctx, arg)
	if err != nil {
		s.logger.Error("CreatePrice sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.CreatePriceResponse{Price: toPrice(result)}, nil
}

func (s *Service) CreateReport(ctx context.Context, req *pb.CreateReportRequest) (*pb.CreateReportResponse, error) {
	var arg CreateReportParams
	if v := req.GetTradingDate(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid TradingDate: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
		arg.TradingDate = v.AsTime()
	} else {
		err := fmt.Errorf("field TradingDate is required%w", validation.ErrUserInput)
		return nil, err
	}
	arg.ItemID = req.GetItemId()
	arg.CurrencyID = req.GetCurrencyId()
	if v := req.GetFirstPrice(); v != nil {
		arg.FirstPrice = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetLastPrice(); v != nil {
		arg.LastPrice = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetMinPrice(); v != nil {
		arg.MinPrice = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetMaxPrice(); v != nil {
		arg.MaxPrice = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetAvgPrice(); v != nil {
		arg.AvgPrice = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetTotalAmount(); v != nil {
		arg.TotalAmount = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetQuantity(); v != nil {
		arg.Quantity = sql.NullString{Valid: true, String: v.Value}
	}

	result, err := s.querier.CreateReport(ctx, arg)
	if err != nil {
		s.logger.Error("CreateReport sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.CreateReportResponse{Report: toReport(result)}, nil
}

func (s *Service) CreateTrade(ctx context.Context, req *pb.CreateTradeRequest) (*pb.CreateTradeResponse, error) {
	var arg CreateTradeParams
	arg.ItemID = req.GetItemId()
	if v := req.GetSellerId(); v != nil {
		arg.SellerID = sql.NullInt32{Valid: true, Int32: v.Value}
	}
	arg.BuyerID = req.GetBuyerId()
	arg.Quantity = req.GetQuantity()
	arg.UnitPrice = req.GetUnitPrice()
	arg.Description = req.GetDescription()
	arg.OfferID = req.GetOfferId()

	result, err := s.querier.CreateTrade(ctx, arg)
	if err != nil {
		s.logger.Error("CreateTrade sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.CreateTradeResponse{Trade: toTrade(result)}, nil
}

func (s *Service) CreateTrader(ctx context.Context, req *pb.CreateTraderRequest) (*pb.CreateTraderResponse, error) {
	var arg CreateTraderParams
	arg.FirstName = req.GetFirstName()
	arg.LastName = req.GetLastName()
	arg.UserName = req.GetUserName()
	arg.Password = req.GetPassword()
	arg.Email = req.GetEmail()
	arg.ConfirmationCode = req.GetConfirmationCode()
	if v := req.GetTimeRegistered(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid TimeRegistered: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
		arg.TimeRegistered = v.AsTime()
	} else {
		err := fmt.Errorf("field TimeRegistered is required%w", validation.ErrUserInput)
		return nil, err
	}
	if v := req.GetTimeConfirmed(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid TimeConfirmed: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
		arg.TimeConfirmed = v.AsTime()
	} else {
		err := fmt.Errorf("field TimeConfirmed is required%w", validation.ErrUserInput)
		return nil, err
	}
	if v := req.GetCountryId(); v != nil {
		arg.CountryID = sql.NullInt32{Valid: true, Int32: v.Value}
	}
	if v := req.GetPreferredCurrencyId(); v != nil {
		arg.PreferredCurrencyID = sql.NullInt32{Valid: true, Int32: v.Value}
	}

	result, err := s.querier.CreateTrader(ctx, arg)
	if err != nil {
		s.logger.Error("CreateTrader sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.CreateTraderResponse{Trader: toTrader(result)}, nil
}

func (s *Service) DeleteCountry(ctx context.Context, req *pb.DeleteCountryRequest) (*pb.DeleteCountryResponse, error) {
	id := req.GetId()

	err := s.querier.DeleteCountry(ctx, id)
	if err != nil {
		s.logger.Error("DeleteCountry sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.DeleteCountryResponse{}, nil
}

func (s *Service) DeleteCurrency(ctx context.Context, req *pb.DeleteCurrencyRequest) (*pb.DeleteCurrencyResponse, error) {
	id := req.GetId()

	err := s.querier.DeleteCurrency(ctx, id)
	if err != nil {
		s.logger.Error("DeleteCurrency sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.DeleteCurrencyResponse{}, nil
}

func (s *Service) DeleteCurrencyRate(ctx context.Context, req *pb.DeleteCurrencyRateRequest) (*pb.DeleteCurrencyRateResponse, error) {
	id := req.GetId()

	err := s.querier.DeleteCurrencyRate(ctx, id)
	if err != nil {
		s.logger.Error("DeleteCurrencyRate sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.DeleteCurrencyRateResponse{}, nil
}

func (s *Service) DeleteCurrencyUsed(ctx context.Context, req *pb.DeleteCurrencyUsedRequest) (*pb.DeleteCurrencyUsedResponse, error) {
	id := req.GetId()

	err := s.querier.DeleteCurrencyUsed(ctx, id)
	if err != nil {
		s.logger.Error("DeleteCurrencyUsed sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.DeleteCurrencyUsedResponse{}, nil
}

func (s *Service) DeleteCurrentInventory(ctx context.Context, req *pb.DeleteCurrentInventoryRequest) (*pb.DeleteCurrentInventoryResponse, error) {
	id := req.GetId()

	err := s.querier.DeleteCurrentInventory(ctx, id)
	if err != nil {
		s.logger.Error("DeleteCurrentInventory sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.DeleteCurrentInventoryResponse{}, nil
}

func (s *Service) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	id := req.GetId()

	err := s.querier.DeleteItem(ctx, id)
	if err != nil {
		s.logger.Error("DeleteItem sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.DeleteItemResponse{}, nil
}

func (s *Service) DeleteOffer(ctx context.Context, req *pb.DeleteOfferRequest) (*pb.DeleteOfferResponse, error) {
	id := req.GetId()

	err := s.querier.DeleteOffer(ctx, id)
	if err != nil {
		s.logger.Error("DeleteOffer sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.DeleteOfferResponse{}, nil
}

func (s *Service) DeletePrice(ctx context.Context, req *pb.DeletePriceRequest) (*pb.DeletePriceResponse, error) {
	id := req.GetId()

	err := s.querier.DeletePrice(ctx, id)
	if err != nil {
		s.logger.Error("DeletePrice sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.DeletePriceResponse{}, nil
}

func (s *Service) DeleteReport(ctx context.Context, req *pb.DeleteReportRequest) (*pb.DeleteReportResponse, error) {
	id := req.GetId()

	err := s.querier.DeleteReport(ctx, id)
	if err != nil {
		s.logger.Error("DeleteReport sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.DeleteReportResponse{}, nil
}

func (s *Service) DeleteTrade(ctx context.Context, req *pb.DeleteTradeRequest) (*pb.DeleteTradeResponse, error) {
	id := req.GetId()

	err := s.querier.DeleteTrade(ctx, id)
	if err != nil {
		s.logger.Error("DeleteTrade sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.DeleteTradeResponse{}, nil
}

func (s *Service) DeleteTrader(ctx context.Context, req *pb.DeleteTraderRequest) (*pb.DeleteTraderResponse, error) {
	id := req.GetId()

	err := s.querier.DeleteTrader(ctx, id)
	if err != nil {
		s.logger.Error("DeleteTrader sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.DeleteTraderResponse{}, nil
}

func (s *Service) GetCountry(ctx context.Context, req *pb.GetCountryRequest) (*pb.GetCountryResponse, error) {
	id := req.GetId()

	result, err := s.querier.GetCountry(ctx, id)
	if err != nil {
		s.logger.Error("GetCountry sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.GetCountryResponse{Country: toCountry(result)}, nil
}

func (s *Service) GetCurrency(ctx context.Context, req *pb.GetCurrencyRequest) (*pb.GetCurrencyResponse, error) {
	id := req.GetId()

	result, err := s.querier.GetCurrency(ctx, id)
	if err != nil {
		s.logger.Error("GetCurrency sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.GetCurrencyResponse{Currency: toCurrency(result)}, nil
}

func (s *Service) GetCurrencyRate(ctx context.Context, req *pb.GetCurrencyRateRequest) (*pb.GetCurrencyRateResponse, error) {
	id := req.GetId()

	result, err := s.querier.GetCurrencyRate(ctx, id)
	if err != nil {
		s.logger.Error("GetCurrencyRate sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.GetCurrencyRateResponse{CurrencyRate: toCurrencyRate(result)}, nil
}

func (s *Service) GetCurrencyUsed(ctx context.Context, req *pb.GetCurrencyUsedRequest) (*pb.GetCurrencyUsedResponse, error) {
	id := req.GetId()

	result, err := s.querier.GetCurrencyUsed(ctx, id)
	if err != nil {
		s.logger.Error("GetCurrencyUsed sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.GetCurrencyUsedResponse{CurrencyUsed: toCurrencyUsed(result)}, nil
}

func (s *Service) GetCurrentInventory(ctx context.Context, req *pb.GetCurrentInventoryRequest) (*pb.GetCurrentInventoryResponse, error) {
	id := req.GetId()

	result, err := s.querier.GetCurrentInventory(ctx, id)
	if err != nil {
		s.logger.Error("GetCurrentInventory sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.GetCurrentInventoryResponse{CurrentInventory: toCurrentInventory(result)}, nil
}

func (s *Service) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	id := req.GetId()

	result, err := s.querier.GetItem(ctx, id)
	if err != nil {
		s.logger.Error("GetItem sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.GetItemResponse{Item: toItem(result)}, nil
}

func (s *Service) GetOffer(ctx context.Context, req *pb.GetOfferRequest) (*pb.GetOfferResponse, error) {
	id := req.GetId()

	result, err := s.querier.GetOffer(ctx, id)
	if err != nil {
		s.logger.Error("GetOffer sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.GetOfferResponse{Offer: toOffer(result)}, nil
}

func (s *Service) GetPrice(ctx context.Context, req *pb.GetPriceRequest) (*pb.GetPriceResponse, error) {
	id := req.GetId()

	result, err := s.querier.GetPrice(ctx, id)
	if err != nil {
		s.logger.Error("GetPrice sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.GetPriceResponse{Price: toPrice(result)}, nil
}

func (s *Service) GetReport(ctx context.Context, req *pb.GetReportRequest) (*pb.GetReportResponse, error) {
	id := req.GetId()

	result, err := s.querier.GetReport(ctx, id)
	if err != nil {
		s.logger.Error("GetReport sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.GetReportResponse{Report: toReport(result)}, nil
}

func (s *Service) GetTrade(ctx context.Context, req *pb.GetTradeRequest) (*pb.GetTradeResponse, error) {
	id := req.GetId()

	result, err := s.querier.GetTrade(ctx, id)
	if err != nil {
		s.logger.Error("GetTrade sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.GetTradeResponse{Trade: toTrade(result)}, nil
}

func (s *Service) GetTrader(ctx context.Context, req *pb.GetTraderRequest) (*pb.GetTraderResponse, error) {
	id := req.GetId()

	result, err := s.querier.GetTrader(ctx, id)
	if err != nil {
		s.logger.Error("GetTrader sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.GetTraderResponse{Trader: toTrader(result)}, nil
}

func (s *Service) ListCountries(ctx context.Context, req *pb.ListCountriesRequest) (*pb.ListCountriesResponse, error) {

	result, err := s.querier.ListCountries(ctx)
	if err != nil {
		s.logger.Error("ListCountries sql call failed", zap.Error(err))
		return nil, err
	}
	res := new(pb.ListCountriesResponse)
	for _, r := range result {
		res.List = append(res.List, toCountry(r))
	}
	return res, nil
}

func (s *Service) ListCurrencies(ctx context.Context, req *pb.ListCurrenciesRequest) (*pb.ListCurrenciesResponse, error) {

	result, err := s.querier.ListCurrencies(ctx)
	if err != nil {
		s.logger.Error("ListCurrencies sql call failed", zap.Error(err))
		return nil, err
	}
	res := new(pb.ListCurrenciesResponse)
	for _, r := range result {
		res.List = append(res.List, toCurrency(r))
	}
	return res, nil
}

func (s *Service) ListCurrenciesUsed(ctx context.Context, req *pb.ListCurrenciesUsedRequest) (*pb.ListCurrenciesUsedResponse, error) {

	result, err := s.querier.ListCurrenciesUsed(ctx)
	if err != nil {
		s.logger.Error("ListCurrenciesUsed sql call failed", zap.Error(err))
		return nil, err
	}
	res := new(pb.ListCurrenciesUsedResponse)
	for _, r := range result {
		res.List = append(res.List, toCurrencyUsed(r))
	}
	return res, nil
}

func (s *Service) ListCurrencyRates(ctx context.Context, req *pb.ListCurrencyRatesRequest) (*pb.ListCurrencyRatesResponse, error) {

	result, err := s.querier.ListCurrencyRates(ctx)
	if err != nil {
		s.logger.Error("ListCurrencyRates sql call failed", zap.Error(err))
		return nil, err
	}
	res := new(pb.ListCurrencyRatesResponse)
	for _, r := range result {
		res.List = append(res.List, toCurrencyRate(r))
	}
	return res, nil
}

func (s *Service) ListCurrentInventory(ctx context.Context, req *pb.ListCurrentInventoryRequest) (*pb.ListCurrentInventoryResponse, error) {

	result, err := s.querier.ListCurrentInventory(ctx)
	if err != nil {
		s.logger.Error("ListCurrentInventory sql call failed", zap.Error(err))
		return nil, err
	}
	res := new(pb.ListCurrentInventoryResponse)
	for _, r := range result {
		res.List = append(res.List, toCurrentInventory(r))
	}
	return res, nil
}

func (s *Service) ListItems(ctx context.Context, req *pb.ListItemsRequest) (*pb.ListItemsResponse, error) {

	result, err := s.querier.ListItems(ctx)
	if err != nil {
		s.logger.Error("ListItems sql call failed", zap.Error(err))
		return nil, err
	}
	res := new(pb.ListItemsResponse)
	for _, r := range result {
		res.List = append(res.List, toItem(r))
	}
	return res, nil
}

func (s *Service) ListOffers(ctx context.Context, req *pb.ListOffersRequest) (*pb.ListOffersResponse, error) {

	result, err := s.querier.ListOffers(ctx)
	if err != nil {
		s.logger.Error("ListOffers sql call failed", zap.Error(err))
		return nil, err
	}
	res := new(pb.ListOffersResponse)
	for _, r := range result {
		res.List = append(res.List, toOffer(r))
	}
	return res, nil
}

func (s *Service) ListPrices(ctx context.Context, req *pb.ListPricesRequest) (*pb.ListPricesResponse, error) {

	result, err := s.querier.ListPrices(ctx)
	if err != nil {
		s.logger.Error("ListPrices sql call failed", zap.Error(err))
		return nil, err
	}
	res := new(pb.ListPricesResponse)
	for _, r := range result {
		res.List = append(res.List, toPrice(r))
	}
	return res, nil
}

func (s *Service) ListReports(ctx context.Context, req *pb.ListReportsRequest) (*pb.ListReportsResponse, error) {

	result, err := s.querier.ListReports(ctx)
	if err != nil {
		s.logger.Error("ListReports sql call failed", zap.Error(err))
		return nil, err
	}
	res := new(pb.ListReportsResponse)
	for _, r := range result {
		res.List = append(res.List, toReport(r))
	}
	return res, nil
}

func (s *Service) ListTraders(ctx context.Context, req *pb.ListTradersRequest) (*pb.ListTradersResponse, error) {

	result, err := s.querier.ListTraders(ctx)
	if err != nil {
		s.logger.Error("ListTraders sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.ListTradersResponse{Trader: toTrader(result)}, nil
}

func (s *Service) ListTrades(ctx context.Context, req *pb.ListTradesRequest) (*pb.ListTradesResponse, error) {

	result, err := s.querier.ListTrades(ctx)
	if err != nil {
		s.logger.Error("ListTrades sql call failed", zap.Error(err))
		return nil, err
	}
	res := new(pb.ListTradesResponse)
	for _, r := range result {
		res.List = append(res.List, toTrade(r))
	}
	return res, nil
}

func (s *Service) UpdateCountry(ctx context.Context, req *pb.UpdateCountryRequest) (*pb.UpdateCountryResponse, error) {
	var arg UpdateCountryParams
	arg.ID = req.GetId()
	arg.Name = req.GetName()

	err := s.querier.UpdateCountry(ctx, arg)
	if err != nil {
		s.logger.Error("UpdateCountry sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.UpdateCountryResponse{}, nil
}

func (s *Service) UpdateCurrency(ctx context.Context, req *pb.UpdateCurrencyRequest) (*pb.UpdateCurrencyResponse, error) {
	var arg UpdateCurrencyParams
	arg.ID = req.GetId()
	arg.Name = req.GetName()
	arg.Code = req.GetCode()
	arg.IsActive = req.GetIsActive()
	arg.IsBaseCurrency = req.GetIsBaseCurrency()

	err := s.querier.UpdateCurrency(ctx, arg)
	if err != nil {
		s.logger.Error("UpdateCurrency sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.UpdateCurrencyResponse{}, nil
}

func (s *Service) UpdateCurrencyRate(ctx context.Context, req *pb.UpdateCurrencyRateRequest) (*pb.UpdateCurrencyRateResponse, error) {
	var arg UpdateCurrencyRateParams
	arg.ID = req.GetId()
	arg.CurrencyID = req.GetCurrencyId()
	arg.BaseCurrencyID = req.GetBaseCurrencyId()
	arg.Rate = req.GetRate()
	if v := req.GetTs(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid Ts: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
		arg.Ts = v.AsTime()
	} else {
		err := fmt.Errorf("field Ts is required%w", validation.ErrUserInput)
		return nil, err
	}

	err := s.querier.UpdateCurrencyRate(ctx, arg)
	if err != nil {
		s.logger.Error("UpdateCurrencyRate sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.UpdateCurrencyRateResponse{}, nil
}

func (s *Service) UpdateCurrencyUsed(ctx context.Context, req *pb.UpdateCurrencyUsedRequest) (*pb.UpdateCurrencyUsedResponse, error) {
	var arg UpdateCurrencyUsedParams
	arg.ID = req.GetId()
	arg.CountryID = req.GetCountryId()
	arg.CurrencyID = req.GetCurrencyId()
	if v := req.GetDateFrom(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid DateFrom: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
		arg.DateFrom = v.AsTime()
	} else {
		err := fmt.Errorf("field DateFrom is required%w", validation.ErrUserInput)
		return nil, err
	}
	if v := req.GetDateTo(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid DateTo: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
		if t := v.AsTime(); !t.IsZero() {
			arg.DateTo.Valid = true
			arg.DateTo.Time = t
		}
	}

	err := s.querier.UpdateCurrencyUsed(ctx, arg)
	if err != nil {
		s.logger.Error("UpdateCurrencyUsed sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.UpdateCurrencyUsedResponse{}, nil
}

func (s *Service) UpdateCurrentInventory(ctx context.Context, req *pb.UpdateCurrentInventoryRequest) (*pb.UpdateCurrentInventoryResponse, error) {
	var arg UpdateCurrentInventoryParams
	arg.ID = req.GetId()
	arg.TraderID = req.GetTraderId()
	arg.ItemID = req.GetItemId()
	arg.Quantity = req.GetQuantity()

	err := s.querier.UpdateCurrentInventory(ctx, arg)
	if err != nil {
		s.logger.Error("UpdateCurrentInventory sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.UpdateCurrentInventoryResponse{}, nil
}

func (s *Service) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {
	var arg UpdateItemParams
	arg.ID = req.GetId()
	arg.Code = req.GetCode()
	arg.Name = req.GetName()
	arg.IsActive = req.GetIsActive()
	arg.CurrencyID = req.GetCurrencyId()
	if v := req.GetDetails(); v != nil {
		arg.Details = sql.NullString{Valid: true, String: v.Value}
	}

	err := s.querier.UpdateItem(ctx, arg)
	if err != nil {
		s.logger.Error("UpdateItem sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.UpdateItemResponse{}, nil
}

func (s *Service) UpdateOffer(ctx context.Context, req *pb.UpdateOfferRequest) (*pb.UpdateOfferResponse, error) {
	var arg UpdateOfferParams
	arg.ID = req.GetId()
	arg.TraderID = req.GetTraderId()
	arg.ItemID = req.GetItemId()
	arg.Quantity = req.GetQuantity()
	arg.Buy = req.GetBuy()
	arg.Sell = req.GetSell()
	if v := req.GetPrice(); v != nil {
		arg.Price = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetTs(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid Ts: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
		arg.Ts = v.AsTime()
	} else {
		err := fmt.Errorf("field Ts is required%w", validation.ErrUserInput)
		return nil, err
	}
	arg.IsActive = req.GetIsActive()

	err := s.querier.UpdateOffer(ctx, arg)
	if err != nil {
		s.logger.Error("UpdateOffer sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.UpdateOfferResponse{}, nil
}

func (s *Service) UpdatePrice(ctx context.Context, req *pb.UpdatePriceRequest) (*pb.UpdatePriceResponse, error) {
	var arg UpdatePriceParams
	arg.ItemID = req.GetItemId()
	arg.CurrencyID = req.GetCurrencyId()
	arg.Buy = req.GetBuy()
	arg.Sell = req.GetSell()
	if v := req.GetTs(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid Ts: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
		arg.Ts = v.AsTime()
	} else {
		err := fmt.Errorf("field Ts is required%w", validation.ErrUserInput)
		return nil, err
	}

	err := s.querier.UpdatePrice(ctx, arg)
	if err != nil {
		s.logger.Error("UpdatePrice sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.UpdatePriceResponse{}, nil
}

func (s *Service) UpdateReport(ctx context.Context, req *pb.UpdateReportRequest) (*pb.UpdateReportResponse, error) {
	var arg UpdateReportParams
	arg.ID = req.GetId()
	if v := req.GetTradingDate(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid TradingDate: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
		arg.TradingDate = v.AsTime()
	} else {
		err := fmt.Errorf("field TradingDate is required%w", validation.ErrUserInput)
		return nil, err
	}
	arg.ItemID = req.GetItemId()
	arg.CurrencyID = req.GetCurrencyId()
	if v := req.GetFirstPrice(); v != nil {
		arg.FirstPrice = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetLastPrice(); v != nil {
		arg.LastPrice = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetMinPrice(); v != nil {
		arg.MinPrice = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetMaxPrice(); v != nil {
		arg.MaxPrice = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetAvgPrice(); v != nil {
		arg.AvgPrice = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetTotalAmount(); v != nil {
		arg.TotalAmount = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetQuantity(); v != nil {
		arg.Quantity = sql.NullString{Valid: true, String: v.Value}
	}

	err := s.querier.UpdateReport(ctx, arg)
	if err != nil {
		s.logger.Error("UpdateReport sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.UpdateReportResponse{}, nil
}

func (s *Service) UpdateTrade(ctx context.Context, req *pb.UpdateTradeRequest) (*pb.UpdateTradeResponse, error) {
	var arg UpdateTradeParams
	arg.ID = req.GetId()
	arg.ItemID = req.GetItemId()
	if v := req.GetSellerId(); v != nil {
		arg.SellerID = sql.NullInt32{Valid: true, Int32: v.Value}
	}
	arg.BuyerID = req.GetBuyerId()
	arg.Quantity = req.GetQuantity()
	arg.UnitPrice = req.GetUnitPrice()
	arg.Description = req.GetDescription()
	arg.OfferID = req.GetOfferId()

	err := s.querier.UpdateTrade(ctx, arg)
	if err != nil {
		s.logger.Error("UpdateTrade sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.UpdateTradeResponse{}, nil
}

func (s *Service) UpdateTrader(ctx context.Context, req *pb.UpdateTraderRequest) (*pb.UpdateTraderResponse, error) {
	var arg UpdateTraderParams
	arg.ID = req.GetId()
	arg.FirstName = req.GetFirstName()
	arg.LastName = req.GetLastName()
	arg.UserName = req.GetUserName()
	arg.Password = req.GetPassword()
	arg.Email = req.GetEmail()
	arg.ConfirmationCode = req.GetConfirmationCode()
	if v := req.GetTimeRegistered(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid TimeRegistered: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
		arg.TimeRegistered = v.AsTime()
	} else {
		err := fmt.Errorf("field TimeRegistered is required%w", validation.ErrUserInput)
		return nil, err
	}
	if v := req.GetTimeConfirmed(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid TimeConfirmed: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
		arg.TimeConfirmed = v.AsTime()
	} else {
		err := fmt.Errorf("field TimeConfirmed is required%w", validation.ErrUserInput)
		return nil, err
	}
	if v := req.GetCountryId(); v != nil {
		arg.CountryID = sql.NullInt32{Valid: true, Int32: v.Value}
	}
	if v := req.GetPreferredCurrencyId(); v != nil {
		arg.PreferredCurrencyID = sql.NullInt32{Valid: true, Int32: v.Value}
	}

	err := s.querier.UpdateTrader(ctx, arg)
	if err != nil {
		s.logger.Error("UpdateTrader sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.UpdateTraderResponse{}, nil
}
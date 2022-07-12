// Code generated by sqlc-grpc (https://github.com/walterwanderley/sqlc-grpc). DO NOT EDIT.

package trades

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "strat_trader/api/trades/v1"
)

func toCountry(in Country) *pb.Country {

	out := new(pb.Country)
	out.Id = in.ID
	out.Name = in.Name
	return out
}

func toCurrency(in Currency) *pb.Currency {

	out := new(pb.Currency)
	out.Id = in.ID
	out.Code = in.Code
	out.Name = in.Name
	out.IsActive = in.IsActive
	out.IsBaseCurrency = in.IsBaseCurrency
	return out
}

func toCurrencyRate(in CurrencyRate) *pb.CurrencyRate {

	out := new(pb.CurrencyRate)
	out.Id = in.ID
	out.CurrencyId = in.CurrencyID
	out.BaseCurrencyId = in.BaseCurrencyID
	out.Rate = in.Rate
	out.Ts = timestamppb.New(in.Ts)
	return out
}

func toCurrencyUsed(in CurrencyUsed) *pb.CurrencyUsed {

	out := new(pb.CurrencyUsed)
	out.Id = in.ID
	out.CountryId = in.CountryID
	out.CurrencyId = in.CurrencyID
	out.DateFrom = timestamppb.New(in.DateFrom)
	if in.DateTo.Valid {
		out.DateTo = timestamppb.New(in.DateTo.Time)
	}
	return out
}

func toCurrentInventory(in CurrentInventory) *pb.CurrentInventory {

	out := new(pb.CurrentInventory)
	out.Id = in.ID
	out.TraderId = in.TraderID
	out.ItemId = in.ItemID
	out.Quantity = in.Quantity
	return out
}

func toItem(in Item) *pb.Item {

	out := new(pb.Item)
	out.Id = in.ID
	out.Code = in.Code
	out.Name = in.Name
	out.IsActive = in.IsActive
	out.CurrencyId = in.CurrencyID
	if in.Details.Valid {
		out.Details = wrapperspb.String(in.Details.String)
	}
	return out
}

func toOffer(in Offer) *pb.Offer {

	out := new(pb.Offer)
	out.Id = in.ID
	out.TraderId = in.TraderID
	out.ItemId = in.ItemID
	out.Quantity = in.Quantity
	out.Buy = in.Buy
	out.Sell = in.Sell
	if in.Price.Valid {
		out.Price = wrapperspb.String(in.Price.String)
	}
	out.Ts = timestamppb.New(in.Ts)
	out.IsActive = in.IsActive
	return out
}

func toPrice(in Price) *pb.Price {

	out := new(pb.Price)
	out.Id = in.ID
	out.ItemId = in.ItemID
	out.CurrencyId = in.CurrencyID
	out.Buy = in.Buy
	out.Sell = in.Sell
	out.Ts = timestamppb.New(in.Ts)
	return out
}

func toReport(in Report) *pb.Report {

	out := new(pb.Report)
	out.Id = in.ID
	out.TradingDate = timestamppb.New(in.TradingDate)
	out.ItemId = in.ItemID
	out.CurrencyId = in.CurrencyID
	if in.FirstPrice.Valid {
		out.FirstPrice = wrapperspb.String(in.FirstPrice.String)
	}
	if in.LastPrice.Valid {
		out.LastPrice = wrapperspb.String(in.LastPrice.String)
	}
	if in.MinPrice.Valid {
		out.MinPrice = wrapperspb.String(in.MinPrice.String)
	}
	if in.MaxPrice.Valid {
		out.MaxPrice = wrapperspb.String(in.MaxPrice.String)
	}
	if in.AvgPrice.Valid {
		out.AvgPrice = wrapperspb.String(in.AvgPrice.String)
	}
	if in.TotalAmount.Valid {
		out.TotalAmount = wrapperspb.String(in.TotalAmount.String)
	}
	if in.Quantity.Valid {
		out.Quantity = wrapperspb.String(in.Quantity.String)
	}
	return out
}

func toTrade(in Trade) *pb.Trade {

	out := new(pb.Trade)
	out.Id = in.ID
	out.ItemId = in.ItemID
	if in.SellerID.Valid {
		out.SellerId = wrapperspb.Int32(in.SellerID.Int32)
	}
	out.BuyerId = in.BuyerID
	out.Quantity = in.Quantity
	out.UnitPrice = in.UnitPrice
	out.Description = in.Description
	out.OfferId = in.OfferID
	return out
}

func toTrader(in Trader) *pb.Trader {

	out := new(pb.Trader)
	out.Id = in.ID
	out.FirstName = in.FirstName
	out.LastName = in.LastName
	out.UserName = in.UserName
	out.Password = in.Password
	out.Email = in.Email
	out.ConfirmationCode = in.ConfirmationCode
	out.TimeRegistered = timestamppb.New(in.TimeRegistered)
	out.TimeConfirmed = timestamppb.New(in.TimeConfirmed)
	if in.CountryID.Valid {
		out.CountryId = wrapperspb.Int32(in.CountryID.Int32)
	}
	if in.PreferredCurrencyID.Valid {
		out.PreferredCurrencyId = wrapperspb.Int32(in.PreferredCurrencyID.Int32)
	}
	return out
}

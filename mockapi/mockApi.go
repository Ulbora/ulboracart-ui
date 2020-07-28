package mockapi

import (
	px "github.com/Ulbora/GoProxy"
	api "github.com/Ulbora/Six910API-Go"
	sdbi "github.com/Ulbora/six910-database-interface"
)

/*
 Six910 is a shopping cart and E-commerce system.
 Copyright (C) 2020 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.
 Copyright (C) 2020 Ken Williamson
 All rights reserved.
 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

//MockAPI MockAPI
type MockAPI struct {
	MockUser           *api.UserResponse
	MockUpdateUserResp *api.Response
}

//GetNew GetNew
func (a *MockAPI) GetNew() api.API {
	return a
}

//SetLogLever SetLogLever
func (a *MockAPI) SetLogLever(level int) {

}

//SetStore SetStore
func (a *MockAPI) SetStore(storeName string, localDomain string) {

}

//SetRestURL SetRestURL
func (a *MockAPI) SetRestURL(url string) {

}

//SetAPIKey SetAPIKey
func (a *MockAPI) SetAPIKey(key string) {

}

//OverrideProxy OverrideProxy
func (a *MockAPI) OverrideProxy(proxy px.Proxy) {

}

//SetStoreID SetStoreID
func (a *MockAPI) SetStoreID(sid int64) {
}

//AddAddress AddAddress
func (a *MockAPI) AddAddress(ad *sdbi.Address, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateAddress UpdateAddress
func (a *MockAPI) UpdateAddress(ad *sdbi.Address, headers *api.Headers) *api.Response {
	return nil
}

//GetAddress GetAddress
func (a *MockAPI) GetAddress(id int64, cid int64, headers *api.Headers) *sdbi.Address {
	return nil
}

//GetAddressList GetAddressList
func (a *MockAPI) GetAddressList(cid int64, headers *api.Headers) *[]sdbi.Address {
	return nil
}

//DeleteAddress DeleteAddress
func (a *MockAPI) DeleteAddress(id int64, cid int64, headers *api.Headers) *api.Response {
	return nil
}

//cart

//AddCart AddCart
func (a *MockAPI) AddCart(c *sdbi.Cart, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateCart UpdateCart
func (a *MockAPI) UpdateCart(c *sdbi.Cart, headers *api.Headers) *api.Response {
	return nil
}

//GetCart GetCart
func (a *MockAPI) GetCart(cid int64, headers *api.Headers) *sdbi.Cart {
	return nil
}

//DeleteCart DeleteCart
func (a *MockAPI) DeleteCart(id int64, cid int64, headers *api.Headers) *api.Response {
	return nil
}

//cartItem

//AddCartItem AddCartItem
func (a *MockAPI) AddCartItem(ci *sdbi.CartItem, cid int64, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateCartItem UpdateCartItem
func (a *MockAPI) UpdateCartItem(ci *sdbi.CartItem, cid int64, headers *api.Headers) *api.Response {
	return nil
}

//GetCartItem GetCartItem
func (a *MockAPI) GetCartItem(cid int64, prodID int64, headers *api.Headers) *sdbi.CartItem {
	return nil
}

//GetCartItemList GetCartItemList
func (a *MockAPI) GetCartItemList(cartID int64, cid int64, headers *api.Headers) *[]sdbi.CartItem {
	return nil
}

//DeleteCartItem DeleteCartItem
func (a *MockAPI) DeleteCartItem(id int64, prodID int64, cartID int64, headers *api.Headers) *api.Response {
	return nil
}

//category

//AddCategory AddCategory
func (a *MockAPI) AddCategory(c *sdbi.Category, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateCategory UpdateCategory
func (a *MockAPI) UpdateCategory(c *sdbi.Category, headers *api.Headers) *api.Response {
	return nil
}

//GetCategory GetCategory
func (a *MockAPI) GetCategory(id int64, headers *api.Headers) *sdbi.Category {
	return nil
}

//GetCategoryList GetCategoryList
func (a *MockAPI) GetCategoryList(headers *api.Headers) *[]sdbi.Category {
	return nil
}

//GetSubCategoryList GetSubCategoryList
func (a *MockAPI) GetSubCategoryList(catID int64, headers *api.Headers) *[]sdbi.Category {
	return nil
}

//DeleteCategory DeleteCategory
func (a *MockAPI) DeleteCategory(id int64, headers *api.Headers) *api.Response {
	return nil
}

//AddCustomer AddCustomer
func (a *MockAPI) AddCustomer(c *sdbi.Customer, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateCustomer UpdateCustomer
func (a *MockAPI) UpdateCustomer(c *sdbi.Customer, headers *api.Headers) *api.Response {
	return nil
}

//GetCustomer GetCustomer
func (a *MockAPI) GetCustomer(email string, headers *api.Headers) *sdbi.Customer {
	return nil
}

//GetCustomerID GetCustomerID
func (a *MockAPI) GetCustomerID(id int64, headers *api.Headers) *sdbi.Customer {
	return nil
}

//GetCustomerList GetCustomerList
func (a *MockAPI) GetCustomerList(headers *api.Headers) *[]sdbi.Customer {
	return nil
}

//DeleteCustomer DeleteCustomer
func (a *MockAPI) DeleteCustomer(id int64, headers *api.Headers) *api.Response {
	return nil
}

//dataStoreWriteLock

//AddDataStoreWriteLock AddDataStoreWriteLock
func (a *MockAPI) AddDataStoreWriteLock(w *sdbi.DataStoreWriteLock, headers *api.Headers) *api.Response {
	return nil
}

//UpdateDataStoreWriteLock UpdateDataStoreWriteLock
func (a *MockAPI) UpdateDataStoreWriteLock(w *sdbi.DataStoreWriteLock, headers *api.Headers) *api.Response {
	return nil
}

//GetDataStoreWriteLock GetDataStoreWriteLock
func (a *MockAPI) GetDataStoreWriteLock(dataStore string, headers *api.Headers) *sdbi.DataStoreWriteLock {
	return nil
}

//dataStore

//AddLocalDatastore AddLocalDatastore
func (a *MockAPI) AddLocalDatastore(d *sdbi.LocalDataStore, headers *api.Headers) *api.Response {
	return nil
}

//UpdateLocalDatastore UpdateLocalDatastore
func (a *MockAPI) UpdateLocalDatastore(d *sdbi.LocalDataStore, headers *api.Headers) *api.Response {
	return nil
}

//GetLocalDatastore GetLocalDatastore
func (a *MockAPI) GetLocalDatastore(dataStoreName string, headers *api.Headers) *sdbi.LocalDataStore {
	return nil
}

//distrubutor

//AddDistributor AddDistributor
func (a *MockAPI) AddDistributor(d *sdbi.Distributor, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateDistributor UpdateDistributor
func (a *MockAPI) UpdateDistributor(d *sdbi.Distributor, headers *api.Headers) *api.Response {
	return nil
}

//GetDistributor GetDistributor
func (a *MockAPI) GetDistributor(id int64, headers *api.Headers) *sdbi.Distributor {
	return nil
}

//GetDistributorList GetDistributorList
func (a *MockAPI) GetDistributorList(headers *api.Headers) *[]sdbi.Distributor {
	return nil
}

//DeleteDistributor DeleteDistributor
func (a *MockAPI) DeleteDistributor(id int64, headers *api.Headers) *api.Response {
	return nil
}

//excluded sub region

//AddExcludedSubRegion AddExcludedSubRegion
func (a *MockAPI) AddExcludedSubRegion(e *sdbi.ExcludedSubRegion, headers *api.Headers) *api.ResponseID {
	return nil
}

//GetExcludedSubRegionList GetExcludedSubRegionList
func (a *MockAPI) GetExcludedSubRegionList(regionID int64, headers *api.Headers) *[]sdbi.ExcludedSubRegion {
	return nil
}

//DeleteExcludedSubRegion DeleteExcludedSubRegion
func (a *MockAPI) DeleteExcludedSubRegion(id int64, regionID int64, headers *api.Headers) *api.Response {
	return nil
}

//included sub region

//AddIncludedSubRegion AddIncludedSubRegion
func (a *MockAPI) AddIncludedSubRegion(e *sdbi.IncludedSubRegion, headers *api.Headers) *api.ResponseID {
	return nil
}

//GetIncludedSubRegionList GetIncludedSubRegionList
func (a *MockAPI) GetIncludedSubRegionList(regionID int64, headers *api.Headers) *[]sdbi.IncludedSubRegion {
	return nil
}

//DeleteIncludedSubRegion DeleteIncludedSubRegion
func (a *MockAPI) DeleteIncludedSubRegion(id int64, regionID int64, headers *api.Headers) *api.Response {
	return nil
}

//instances

//AddInstance AddInstance
func (a *MockAPI) AddInstance(i *sdbi.Instances, headers *api.Headers) *api.Response {
	return nil
}

//UpdateInstance UpdateInstance
func (a *MockAPI) UpdateInstance(i *sdbi.Instances, headers *api.Headers) *api.Response {
	return nil
}

//GetInstance GetInstance
func (a *MockAPI) GetInstance(name string, dataStoreName string, headers *api.Headers) *sdbi.Instances {
	return nil
}

//GetInstanceList GetInstanceList
func (a *MockAPI) GetInstanceList(dataStoreName string, headers *api.Headers) *[]sdbi.Instances {
	return nil
}

// insurance

//AddInsurance AddInsurance
func (a *MockAPI) AddInsurance(i *sdbi.Insurance, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateInsurance UpdateInsurance
func (a *MockAPI) UpdateInsurance(i *sdbi.Insurance, headers *api.Headers) *api.Response {
	return nil
}

//GetInsurance GetInsurance
func (a *MockAPI) GetInsurance(id int64, headers *api.Headers) *sdbi.Insurance {
	return nil
}

//GetInsuranceList GetInsuranceList
func (a *MockAPI) GetInsuranceList(headers *api.Headers) *[]sdbi.Insurance {
	return nil
}

//DeleteInsurance DeleteInsurance
func (a *MockAPI) DeleteInsurance(id int64, headers *api.Headers) *api.Response {
	return nil
}

//order

//AddOrder AddOrder
func (a *MockAPI) AddOrder(o *sdbi.Order, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateOrder UpdateOrder
func (a *MockAPI) UpdateOrder(o *sdbi.Order, headers *api.Headers) *api.Response {
	return nil
}

//GetOrder GetOrder
func (a *MockAPI) GetOrder(id int64, headers *api.Headers) *sdbi.Order {
	return nil
}

//GetOrderList GetOrderList
func (a *MockAPI) GetOrderList(cid int64, headers *api.Headers) *[]sdbi.Order {
	return nil
}

//DeleteOrder DeleteOrder
func (a *MockAPI) DeleteOrder(id int64, headers *api.Headers) *api.Response {
	return nil
}

//order comments

//AddOrderComments AddOrderComments
func (a *MockAPI) AddOrderComments(c *sdbi.OrderComment, headers *api.Headers) *api.ResponseID {
	return nil
}

//GetOrderCommentList GetOrderCommentList
func (a *MockAPI) GetOrderCommentList(orderID int64, headers *api.Headers) *[]sdbi.OrderComment {
	return nil
}

//order items

//AddOrderItem AddOrderItem
func (a *MockAPI) AddOrderItem(i *sdbi.OrderItem, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateOrderItem UpdateOrderItem
func (a *MockAPI) UpdateOrderItem(i *sdbi.OrderItem, headers *api.Headers) *api.Response {
	return nil
}

//GetOrderItem GetOrderItem
func (a *MockAPI) GetOrderItem(id int64, headers *api.Headers) *sdbi.OrderItem {
	return nil
}

//GetOrderItemList GetOrderItemList
func (a *MockAPI) GetOrderItemList(orderID int64, headers *api.Headers) *[]sdbi.OrderItem {
	return nil
}

//DeleteOrderItem DeleteOrderItem
func (a *MockAPI) DeleteOrderItem(id int64, headers *api.Headers) *api.Response {
	return nil
}

//order transaction

//AddOrderTransaction AddOrderTransaction
func (a *MockAPI) AddOrderTransaction(t *sdbi.OrderTransaction, headers *api.Headers) *api.ResponseID {
	return nil
}

//GetOrderTransactionList GetOrderTransactionList
func (a *MockAPI) GetOrderTransactionList(orderID int64, headers *api.Headers) *[]sdbi.OrderTransaction {
	return nil
}

//payment gateway

//AddPaymentGateway AddPaymentGateway
func (a *MockAPI) AddPaymentGateway(pgw *sdbi.PaymentGateway, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdatePaymentGateway UpdatePaymentGateway
func (a *MockAPI) UpdatePaymentGateway(pgw *sdbi.PaymentGateway, headers *api.Headers) *api.Response {
	return nil
}

//GetPaymentGateway GetPaymentGateway
func (a *MockAPI) GetPaymentGateway(id int64, headers *api.Headers) *sdbi.PaymentGateway {
	return nil
}

//GetPaymentGateways GetPaymentGateways
func (a *MockAPI) GetPaymentGateways(headers *api.Headers) *[]sdbi.PaymentGateway {
	return nil
}

//DeletePaymentGateway DeletePaymentGateway
func (a *MockAPI) DeletePaymentGateway(id int64, headers *api.Headers) *api.Response {
	return nil
}

//plugins

//AddPlugin AddPlugin
func (a *MockAPI) AddPlugin(p *sdbi.Plugins, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdatePlugin UpdatePlugin
func (a *MockAPI) UpdatePlugin(p *sdbi.Plugins, headers *api.Headers) *api.Response {
	return nil
}

//GetPlugin GetPlugin
func (a *MockAPI) GetPlugin(id int64, headers *api.Headers) *sdbi.Plugins {
	return nil
}

//GetPluginList GetPluginList
func (a *MockAPI) GetPluginList(start int64, end int64, headers *api.Headers) *[]sdbi.Plugins {
	return nil
}

//DeletePlugin DeletePlugin
func (a *MockAPI) DeletePlugin(id int64, headers *api.Headers) *api.Response {
	return nil
}

//products

//AddProduct AddProduct
func (a *MockAPI) AddProduct(p *sdbi.Product, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateProduct UpdateProduct
func (a *MockAPI) UpdateProduct(p *sdbi.Product, headers *api.Headers) *api.Response {
	return nil
}

//GetProductByID GetProductByID
func (a *MockAPI) GetProductByID(id int64, headers *api.Headers) *sdbi.Product {
	return nil
}

//GetProductsByName GetProductsByName
func (a *MockAPI) GetProductsByName(name string, start int64, end int64, headers *api.Headers) *[]sdbi.Product {
	return nil
}

//GetProductsByCaterory GetProductsByCaterory
func (a *MockAPI) GetProductsByCaterory(catID int64, start int64, end int64, headers *api.Headers) *[]sdbi.Product {
	return nil
}

//GetProductList GetProductList
func (a *MockAPI) GetProductList(start int64, end int64, headers *api.Headers) *[]sdbi.Product {
	return nil
}

//DeleteProduct DeleteProduct
func (a *MockAPI) DeleteProduct(id int64, headers *api.Headers) *api.Response {
	return nil
}

//product category

//AddProductCategory AddProductCategory
func (a *MockAPI) AddProductCategory(pc *sdbi.ProductCategory, headers *api.Headers) *api.Response {
	return nil
}

//DeleteProductCategory DeleteProductCategory
func (a *MockAPI) DeleteProductCategory(pc *sdbi.ProductCategory, headers *api.Headers) *api.Response {
	return nil
}

//region

//AddRegion AddRegion
func (a *MockAPI) AddRegion(r *sdbi.Region, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateRegion UpdateRegion
func (a *MockAPI) UpdateRegion(r *sdbi.Region, headers *api.Headers) *api.Response {
	return nil
}

//GetRegion GetRegion
func (a *MockAPI) GetRegion(id int64, headers *api.Headers) *sdbi.Region {
	return nil
}

//GetRegionList GetRegionList
func (a *MockAPI) GetRegionList(headers *api.Headers) *[]sdbi.Region {
	return nil
}

//DeleteRegion DeleteRegion
func (a *MockAPI) DeleteRegion(id int64, headers *api.Headers) *api.Response {
	return nil
}

//shipment

//AddShipment AddShipment
func (a *MockAPI) AddShipment(s *sdbi.Shipment, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateShipment UpdateShipment
func (a *MockAPI) UpdateShipment(s *sdbi.Shipment, headers *api.Headers) *api.Response {
	return nil
}

//GetShipment GetShipment
func (a *MockAPI) GetShipment(id int64, headers *api.Headers) *sdbi.Shipment {
	return nil
}

//GetShipmentList GetShipmentList
func (a *MockAPI) GetShipmentList(orderID int64, headers *api.Headers) *[]sdbi.Shipment {
	return nil
}

//DeleteShipment DeleteShipment
func (a *MockAPI) DeleteShipment(id int64, headers *api.Headers) *api.Response {
	return nil
}

//shipment box

//AddShipmentBox AddShipmentBox
func (a *MockAPI) AddShipmentBox(sb *sdbi.ShipmentBox, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateShipmentBox UpdateShipmentBox
func (a *MockAPI) UpdateShipmentBox(sb *sdbi.ShipmentBox, headers *api.Headers) *api.Response {
	return nil
}

//GetShipmentBox GetShipmentBox
func (a *MockAPI) GetShipmentBox(id int64, headers *api.Headers) *sdbi.ShipmentBox {
	return nil
}

//GetShipmentBoxList GetShipmentBoxList
func (a *MockAPI) GetShipmentBoxList(shipmentID int64, headers *api.Headers) *[]sdbi.ShipmentBox {
	return nil
}

//DeleteShipmentBox DeleteShipmentBox
func (a *MockAPI) DeleteShipmentBox(id int64, headers *api.Headers) *api.Response {
	return nil
}

//shipment item

//AddShipmentItem AddShipmentItem
func (a *MockAPI) AddShipmentItem(si *sdbi.ShipmentItem, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateShipmentItem UpdateShipmentItem
func (a *MockAPI) UpdateShipmentItem(si *sdbi.ShipmentItem, headers *api.Headers) *api.Response {
	return nil
}

//GetShipmentItem GetShipmentItem
func (a *MockAPI) GetShipmentItem(id int64, headers *api.Headers) *sdbi.ShipmentItem {
	return nil
}

//GetShipmentItemList GetShipmentItemList
func (a *MockAPI) GetShipmentItemList(shipmentID int64, headers *api.Headers) *[]sdbi.ShipmentItem {
	return nil
}

//GetShipmentItemListByBox GetShipmentItemListByBox
func (a *MockAPI) GetShipmentItemListByBox(boxNumber int64, shipmentID int64, headers *api.Headers) *[]sdbi.ShipmentItem {
	return nil
}

//DeleteShipmentItem DeleteShipmentItem
func (a *MockAPI) DeleteShipmentItem(id int64, headers *api.Headers) *api.Response {
	return nil
}

//shipment carrier

//AddShippingCarrier AddShippingCarrier
func (a *MockAPI) AddShippingCarrier(c *sdbi.ShippingCarrier, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateShippingCarrier UpdateShippingCarrier
func (a *MockAPI) UpdateShippingCarrier(c *sdbi.ShippingCarrier, headers *api.Headers) *api.Response {
	return nil
}

//GetShippingCarrier GetShippingCarrier
func (a *MockAPI) GetShippingCarrier(id int64, headers *api.Headers) *sdbi.ShippingCarrier {
	return nil
}

//GetShippingCarrierList GetShippingCarrierList
func (a *MockAPI) GetShippingCarrierList(headers *api.Headers) *[]sdbi.ShippingCarrier {
	return nil
}

//DeleteShippingCarrier DeleteShippingCarrier
func (a *MockAPI) DeleteShippingCarrier(id int64, headers *api.Headers) *api.Response {
	return nil
}

//shipment method

//AddShippingMethod AddShippingMethod
func (a *MockAPI) AddShippingMethod(s *sdbi.ShippingMethod, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateShippingMethod UpdateShippingMethod
func (a *MockAPI) UpdateShippingMethod(s *sdbi.ShippingMethod, headers *api.Headers) *api.Response {
	return nil
}

//GetShippingMethod GetShippingMethod
func (a *MockAPI) GetShippingMethod(id int64, headers *api.Headers) *sdbi.ShippingMethod {
	return nil
}

//GetShippingMethodList GetShippingMethodList
func (a *MockAPI) GetShippingMethodList(headers *api.Headers) *[]sdbi.ShippingMethod {
	return nil
}

//DeleteShippingMethod DeleteShippingMethod
func (a *MockAPI) DeleteShippingMethod(id int64, headers *api.Headers) *api.Response {
	return nil
}

//store

//AddStore AddStore
func (a *MockAPI) AddStore(s *sdbi.Store, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateStore UpdateStore
func (a *MockAPI) UpdateStore(s *sdbi.Store, headers *api.Headers) *api.Response {
	return nil
}

//GetStore GetStore
func (a *MockAPI) GetStore(sname string, localDomain string, headers *api.Headers) *sdbi.Store {
	return nil
}

//DeleteStore DeleteStore
func (a *MockAPI) DeleteStore(sname string, localDomain string, headers *api.Headers) *api.Response {
	return nil
}

//store plugin

//AddStorePlugin AddStorePlugin
func (a *MockAPI) AddStorePlugin(sp *sdbi.StorePlugins, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateStorePlugin UpdateStorePlugin
func (a *MockAPI) UpdateStorePlugin(sp *sdbi.StorePlugins, headers *api.Headers) *api.Response {
	return nil
}

//GetStorePlugin GetStorePlugin
func (a *MockAPI) GetStorePlugin(id int64, headers *api.Headers) *sdbi.StorePlugins {
	return nil
}

//GetStorePluginList GetStorePluginList
func (a *MockAPI) GetStorePluginList(headers *api.Headers) *[]sdbi.StorePlugins {
	return nil
}

//DeleteStorePlugin DeleteStorePlugin
func (a *MockAPI) DeleteStorePlugin(id int64, headers *api.Headers) *api.Response {
	return nil
}

//sub region

//AddSubRegion AddSubRegion
func (a *MockAPI) AddSubRegion(s *sdbi.SubRegion, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateSubRegion UpdateSubRegion
func (a *MockAPI) UpdateSubRegion(s *sdbi.SubRegion, headers *api.Headers) *api.Response {
	return nil
}

//GetSubRegion GetSubRegion
func (a *MockAPI) GetSubRegion(id int64, headers *api.Headers) *sdbi.SubRegion {
	return nil
}

//GetSubRegionList GetSubRegionList
func (a *MockAPI) GetSubRegionList(regionID int64, headers *api.Headers) *[]sdbi.SubRegion {
	return nil
}

//DeleteSubRegion DeleteSubRegion
func (a *MockAPI) DeleteSubRegion(id int64, headers *api.Headers) *api.Response {
	return nil
}

//user

//AddCustomerUser AddCustomerUser
func (a *MockAPI) AddCustomerUser(u *api.User, headers *api.Headers) *api.Response {
	return nil
}

//UpdateUser UpdateUser
func (a *MockAPI) UpdateUser(u *api.User, headers *api.Headers) *api.Response {
	return a.MockUpdateUserResp
}

//GetUser GetUser
func (a *MockAPI) GetUser(u *api.User, headers *api.Headers) *api.UserResponse {
	return a.MockUser
}

//GetAdminUsers GetAdminUsers
func (a *MockAPI) GetAdminUsers(headers *api.Headers) *[]api.UserResponse {
	return nil
}

//GetCustomerUsers GetCustomerUsers
func (a *MockAPI) GetCustomerUsers(headers *api.Headers) *[]api.UserResponse {
	return nil
}

//zip code zone

//AddZoneZip AddZoneZip
func (a *MockAPI) AddZoneZip(z *sdbi.ZoneZip, headers *api.Headers) *api.ResponseID {
	return nil
}

//GetZoneZipListByExclusion GetZoneZipListByExclusion
func (a *MockAPI) GetZoneZipListByExclusion(exID int64, headers *api.Headers) *[]sdbi.ZoneZip {
	return nil
}

//GetZoneZipListByInclusion GetZoneZipListByInclusion
func (a *MockAPI) GetZoneZipListByInclusion(incID int64, headers *api.Headers) *[]sdbi.ZoneZip {
	return nil
}

//DeleteZoneZip DeleteZoneZip
func (a *MockAPI) DeleteZoneZip(id int64, incID int64, exID int64, headers *api.Headers) *api.Response {
	return nil
}
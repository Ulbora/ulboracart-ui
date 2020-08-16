package handlers

import (
	"net/http"
	"strconv"
	"sync"

	api "github.com/Ulbora/Six910API-Go"
	six910api "github.com/Ulbora/Six910API-Go"
	sdbi "github.com/Ulbora/six910-database-interface"
	"github.com/gorilla/mux"
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

//ShipPage ShipPage
type ShipPage struct {
	Error         string
	Shipment      *sdbi.Shipment
	ShipmentItems *[]sdbi.ShipmentItem
	ShipmentBoxes *[]sdbi.ShipmentBox
	Shipments     *[]sdbi.Shipment
	Order         *sdbi.Order
	OrderItems    *[]sdbi.OrderItem
	OrderComments *[]sdbi.OrderComment
}

//StoreAdminAddShipmentPage StoreAdminAddShipmentPage
func (h *Six910Handler) StoreAdminAddShipmentPage(w http.ResponseWriter, r *http.Request) {
	gss, suc := h.getSession(r)
	h.Log.Debug("session suc in shipment add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gss) {
			asvars := mux.Vars(r)
			asidstr := asvars["id"]
			asOIID, _ := strconv.ParseInt(asidstr, 10, 64)
			aspErr := r.URL.Query().Get("error")
			var page ShipPage
			page.Error = aspErr
			hd := h.getHeader(gss)
			var wg sync.WaitGroup
			wg.Add(1)
			go func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				page.Order = h.API.GetOrder(oid, header)
			}(asOIID, hd)

			wg.Add(1)
			go func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				page.OrderComments = h.API.GetOrderCommentList(oid, header)
			}(asOIID, hd)

			wg.Add(1)
			go func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				page.OrderItems = h.API.GetOrderItemList(oid, header)
			}(asOIID, hd)

			wg.Wait()
			h.Log.Debug("shipment page", page)
			// h.Log.Debug("shipment order", *page.Order)
			// h.Log.Debug("shipment order notes", *page.OrderComments)
			// h.Log.Debug("shipment order items", *page.OrderItems)

			h.AdminTemplates.ExecuteTemplate(w, adminAddShipmentPage, &page)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminAddShipment StoreAdminAddShipment
func (h *Six910Handler) StoreAdminAddShipment(w http.ResponseWriter, r *http.Request) {
	as, suc := h.getSession(r)
	h.Log.Debug("session suc in shipment add", suc)
	if suc {
		if h.isStoreAdminLoggedIn(as) {
			sh := h.processShipment(r)
			h.Log.Debug("shipment in add", *sh)
			hd := h.getHeader(as)
			shres := h.API.AddShipment(sh, hd)
			h.Log.Debug("shipment add resp", *shres)
			var success = true
			if shres.Success {
				oil := h.API.GetOrderItemList(sh.OrderID, hd)
				var oichan = make(chan *api.ResponseID, len(*oil))
				var wg sync.WaitGroup
				for i := range *oil {
					wg.Add(1)
					go func(oi *sdbi.OrderItem, header *six910api.Headers, ch chan *api.ResponseID) {
						defer wg.Done()
						h.Log.Debug("order item in goroutine", *oi)
						var si sdbi.ShipmentItem
						si.OrderItemID = oi.ID
						si.Quantity = oi.Quantity
						si.ShipmentID = shres.ID
						h.Log.Debug("shipment item in goroutine", si)
						ires := h.API.AddShipmentItem(&si, header)
						ch <- ires
					}(&(*oil)[i], hd, oichan)
				}
				wg.Wait()
				close(oichan)
				for res := range oichan {
					if !res.Success {
						success = false
					}
				}
			} else {
				success = false
			}
			h.Log.Debug("shipment all add suc", success)
			if success {
				http.Redirect(w, r, adminOrderListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminAddShipmentViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditShipmentPage StoreAdminEditShipmentPage
func (h *Six910Handler) StoreAdminEditShipmentPage(w http.ResponseWriter, r *http.Request) {
	ess, suc := h.getSession(r)
	h.Log.Debug("session suc in shipment edit view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(ess) {
			var esparm ShipPage
			edErr := r.URL.Query().Get("error")
			esparm.Error = edErr

			hd := h.getHeader(ess)
			esvars := mux.Vars(r)
			esidstr := esvars["id"]
			esID, _ := strconv.ParseInt(esidstr, 10, 64)
			h.Log.Debug("shipment id in edit", esID)

			ship := h.API.GetShipment(esID, hd)
			esparm.Shipment = ship
			h.Log.Debug("shipment in edit", ship)

			var wg sync.WaitGroup
			wg.Add(1)
			go func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				esparm.Order = h.API.GetOrder(oid, header)
			}(ship.OrderID, hd)

			wg.Add(1)
			go func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				esparm.OrderComments = h.API.GetOrderCommentList(oid, header)
			}(ship.OrderID, hd)

			wg.Add(1)
			go func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				esparm.OrderItems = h.API.GetOrderItemList(oid, header)
			}(ship.OrderID, hd)

			wg.Add(1)
			go func(spid int64, header *six910api.Headers) {
				defer wg.Done()
				esparm.ShipmentBoxes = h.API.GetShipmentBoxList(spid, header)
			}(ship.ID, hd)

			wg.Add(1)
			go func(spid int64, header *six910api.Headers) {
				defer wg.Done()
				esparm.ShipmentItems = h.API.GetShipmentItemList(spid, header)
			}(ship.ID, hd)

			wg.Wait()

			h.Log.Debug("shipment page", esparm)

			h.AdminTemplates.ExecuteTemplate(w, adminEditShipmentPage, &esparm)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processShipment(r *http.Request) *sdbi.Shipment {
	var p sdbi.Shipment
	id := r.FormValue("id")
	p.ID, _ = strconv.ParseInt(id, 10, 64)
	p.Status = r.FormValue("status")
	boxes := r.FormValue("boxes")
	p.Boxes, _ = strconv.ParseInt(boxes, 10, 64)
	shippingHandling := r.FormValue("shippingHandling")
	p.ShippingHandling, _ = strconv.ParseFloat(shippingHandling, 64)
	insurance := r.FormValue("insurance")
	p.Insurance, _ = strconv.ParseFloat(insurance, 64)
	orderID := r.FormValue("orderId")
	p.OrderID, _ = strconv.ParseInt(orderID, 10, 64)

	return &p
}

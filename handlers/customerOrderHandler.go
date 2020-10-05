package handlers

import (
	"net/http"
	"strconv"
	"sync"

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

//OrderViewPage OrderViewPage
type OrderViewPage struct {
	Order    *sdbi.Order
	Items    *[]sdbi.OrderItem
	Comments *[]sdbi.OrderComment
}

//ViewCustomerOrder ViewCustomerOrder
func (h *Six910Handler) ViewCustomerOrder(w http.ResponseWriter, r *http.Request) {
	vorps, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreCustomerLoggedIn(vorps) {
			vodrvars := mux.Vars(r)
			vodridstr := vodrvars["id"]
			vodrid, _ := strconv.ParseInt(vodridstr, 10, 64)

			h.Log.Debug("vodrid: ", vodrid)
			hd := h.getHeader(vorps)
			var ovpage OrderViewPage

			var wg sync.WaitGroup

			wg.Add(1)
			func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				ovpage.Order = h.API.GetOrder(oid, header)
			}(vodrid, hd)

			wg.Add(1)
			func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				ovpage.Items = h.API.GetOrderItemList(oid, header)
			}(vodrid, hd)

			wg.Add(1)
			func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				ovpage.Comments = h.API.GetOrderCommentList(oid, header)
			}(vodrid, hd)

			wg.Wait()

			h.Log.Debug("ovpage: ", ovpage)
			h.Templates.ExecuteTemplate(w, customerOrderPage, &ovpage)
		} else {
			http.Redirect(w, r, customerLoginView, http.StatusFound)
		}
	}
}

//ViewCustomerOrderList ViewCustomerOrderList
func (h *Six910Handler) ViewCustomerOrderList(w http.ResponseWriter, r *http.Request) {
	vorrls, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreCustomerLoggedIn(vorrls) {
			var cid int64
			var cidi int
			fcid := vorrls.Values["customerId"]
			if fcid != nil {
				cidi = fcid.(int)
				cid = int64(cidi)
			}
			h.Log.Debug("cid: ", cid)
			hd := h.getHeader(vorrls)

			odlst := h.API.GetOrderList(cid, hd)

			h.Log.Debug("odlst: ", odlst)
			h.Templates.ExecuteTemplate(w, customerOrderListPage, &odlst)

		} else {
			http.Redirect(w, r, customerLoginView, http.StatusFound)
		}
	}
}
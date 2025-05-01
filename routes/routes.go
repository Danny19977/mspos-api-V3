package routes

import (
	"github.com/danny19977/mspos-api-v3/controllers/area"
	"github.com/danny19977/mspos-api-v3/controllers/asm"
	"github.com/danny19977/mspos-api-v3/controllers/auth"
	"github.com/danny19977/mspos-api-v3/controllers/brand"
	"github.com/danny19977/mspos-api-v3/controllers/commune"
	"github.com/danny19977/mspos-api-v3/controllers/country"
	"github.com/danny19977/mspos-api-v3/controllers/cyclo"
	"github.com/danny19977/mspos-api-v3/controllers/dashboard"
	"github.com/danny19977/mspos-api-v3/controllers/dr"
	"github.com/danny19977/mspos-api-v3/controllers/manager"
	"github.com/danny19977/mspos-api-v3/controllers/pos"
	"github.com/danny19977/mspos-api-v3/controllers/posequiment"
	"github.com/danny19977/mspos-api-v3/controllers/posform"
	PosFormItem "github.com/danny19977/mspos-api-v3/controllers/posformitem"
	"github.com/danny19977/mspos-api-v3/controllers/province"
	"github.com/danny19977/mspos-api-v3/controllers/routeplan.go"
	RoutePlanItem "github.com/danny19977/mspos-api-v3/controllers/routeplanitem"
	Subarea "github.com/danny19977/mspos-api-v3/controllers/subarea"
	"github.com/danny19977/mspos-api-v3/controllers/sup"
	"github.com/danny19977/mspos-api-v3/controllers/user"
	"github.com/danny19977/mspos-api-v3/controllers/user_logs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {

	api := app.Group("/api", logger.New())

	// Authentification controller
	a := api.Group("/auth")
	a.Post("/register", auth.Register)
	a.Post("/login", auth.Login)
	a.Post("/forgot-password", auth.Forgot)
	a.Post("/reset/:token", auth.ResetPassword)

	// app.Use(middlewares.IsAuthenticated)

	a.Get("/user", auth.AuthUser)
	a.Put("/profil/info", auth.UpdateInfo)
	a.Put("/change-password", auth.ChangePassword)
	a.Post("/logout", auth.Logout)

	// Users controller
	u := api.Group("/users")
	u.Get("/all", user.GetAllUsers)
	u.Get("/all/paginate", user.GetPaginatedUsers)
	u.Get("/all/paginate/nosearch", user.GetPaginatedNoSerach)
	u.Get("/all/:id", user.GetUserByID)
	u.Get("/get/:uuid", user.GetUser)
	u.Post("/create", user.CreateUser)
	u.Put("/update/:uuid", user.UpdateUser)
	u.Delete("/delete/:uuid", user.DeleteUser)

	// Province controller
	prov := api.Group("/provinces")
	prov.Get("/all", province.GetAllProvinces)
	prov.Get("/all/:country_uuid", province.GetAllProvinceByCountry)
	prov.Get("/all/paginate", province.GetPaginatedProvince)
	prov.Get("/all/paginate/province/:province_uuid", province.GetPaginatedASM)
	prov.Get("/all/:id", province.GetProvinceByID)
	prov.Get("/get/:uuid", province.GetProvince)
	prov.Get("/get-by/:name", province.GetProvinceByName)
	prov.Post("/create", province.CreateProvince)
	prov.Put("/update/:uuid", province.UpdateProvince)
	prov.Delete("/delete/:uuid", province.DeleteProvince)

	// Areas controller
	ar := api.Group("/areas")
	ar.Get("/all", area.GetAllAreas)
	ar.Get("/all/province_uuid", area.GetAllAreasByProvinceUUID)
	ar.Get("/all/paginate", area.GetPaginatedAreas)
	ar.Get("/all/paginate/province/:province_uuid", area.GetAreaByASM)
	ar.Get("/all/paginate/area/:area_uuid", area.GetAreaBySups)
	ar.Get("/all/:id", area.GetAreaByID)
	ar.Get("/all-area/:id", area.GetSupAreaByID)
	ar.Get("/get/:uuid", area.GetArea)
	ar.Get("/get-by/:name", area.GetAreaByName)
	ar.Post("/create", area.CreateArea)
	ar.Put("/update/:uuid", area.UpdateArea)
	ar.Delete("/delete/:uuid", area.DeleteArea)

	//SubArea controller
	sa := api.Group("/subareas")
	sa.Get("/all", Subarea.GetAllSubArea)
	sa.Get("/all/:area_uuid", Subarea.GetAllDataBySubAreaByAreaUUID)
	sa.Get("/all/paginate", Subarea.GetPaginatedSubArea)
	sa.Get("/all/paginate/province/:province_uuid", Subarea.GetPaginatedSubAreaByASM)
	sa.Get("/all/paginate/area/:area_uuid", Subarea.GetPaginatedSubAreaBySup)
	sa.Get("/all/paginate/subarea/:subarea_uuid", Subarea.GetAllSubAreaDr)
	sa.Get("/get/:uuid", Subarea.GetSubArea)
	sa.Get("/get-by/:name", Subarea.GetSubAreaByName)
	sa.Post("/create", Subarea.CreateSubArea)
	sa.Put("/update/:uuid", Subarea.UpdateSubArea)
	sa.Delete("/delete/:uuid", Subarea.DeleteSubarea)

	// Commune controller
	com := api.Group("/communes")
	com.Get("/all", commune.GetAllCommunes)
	com.Get("/all/:subarea_uuid", commune.GetAllCommunesBySubAreaUUID)
	com.Get("/all/paginate", commune.GetPaginatedCommunes)
	com.Get("/all/paginate/province/:province_uuid", commune.GetPaginatedCommunesByProvinceUUID)
	com.Get("/all/paginate/area/:area_uuid", commune.GetPaginatedCommunesByAreaUUID)
	com.Get("/all/paginate/subarea/:subarea_uuid", commune.GetPaginatedCommunesBySubAreaUUID)
	com.Get("/all/paginate/commune/:commune_uuid", commune.GetPaginatedCommunesByCyclo)
	com.Get("/all/:id", commune.GetCommune)
	com.Get("/get/:uuid", commune.GetCommune)
	com.Post("/create", commune.CreateCommune)
	com.Put("/update/:uuid", commune.UpdateCommune)
	com.Delete("/delete/:uuid", commune.DeleteCommune)

	// ASM controller
	as := api.Group("/asms")
	as.Get("/all", asm.GetAllAsms)
	as.Get("/all/paginate", asm.GetPaginatedASM)
	as.Get("/all/paginate/province/:province_uuid", asm.GetPaginatedASMByProvince)
	// as.Get("/all/:id", asm.GetAsmByID)
	as.Post("/create", asm.CreateAsm)
	as.Get("/get/:uuid", asm.GetAsm)
	as.Put("/update/:uuid", asm.UpdateAsm)
	as.Delete("/delete/:uuid", asm.DeleteAsm)

	// Manager controller
	ma := api.Group("/managers")
	ma.Get("/all", manager.GetAllManagers)
	ma.Get("/all/paginate", manager.GetPaginatedManager)
	ma.Get("/get/:uuid", manager.GetManager)
	// ma.Get("/all/:id", manager.GetManagerByID)
	ma.Post("/create", manager.CreateManager)
	ma.Put("/update/:uuid", manager.UpdateManager)
	ma.Delete("/delete/:uuid", manager.DeleteManager)

	// Posforms controller
	posf := api.Group("/posforms")
	posf.Get("/all", posform.GetAllPosforms)
	posf.Get("/all/paginate", posform.GetPaginatedPosForm)
	posf.Get("/all/paginate/province/:province_uuid", posform.GetPaginatedPosFormProvine)
	posf.Get("/all/paginate/area/:area_uuid", posform.GetPaginatedPosFormArea)
	posf.Get("/all/paginate/subarea/:subarea_uuid", posform.GetPaginatedPosFormSubArea)
	posf.Get("/all/paginate/:cyclo_uuid", posform.GetPaginatedPosFormCommune)
	posf.Get("/all/paginate/:pos_uuid", posform.GetPaginatedPosFormByPOS)
	posf.Post("/create", posform.CreatePosform)
	posf.Get("/get/:uuid", posform.GetPosform)
	posf.Put("/update/:uuid", posform.UpdatePosform)
	posf.Delete("/delete/:uuid", posform.DeletePosform)

	//POSformItem controller
	posfi := api.Group("/posform-items")
	posfi.Get("/all/", PosFormItem.GetAllPosFormItems)
	posfi.Get("/all/:posform_uuid", PosFormItem.GetAllPosFormItemsByUUID)
	posfi.Get("/all/paginate", PosFormItem.GetPaginatedPosformItem)
	// posfi.Get("/get/:uuid", PosFormItem.Get)
	posfi.Post("/create", PosFormItem.CreatePosformItem)
	posfi.Put("/update/:uuid", PosFormItem.UpdatePosformItem)
	posfi.Delete("/delete/:uuid", PosFormItem.DeletePosformItem)

	// Sup controller
	su := api.Group("/sups")
	su.Get("/all", sup.GetAllSups)
	su.Get("/all/paginate", sup.GetPaginatedSups)
	su.Get("/all/paginate/province/:province_uuid", sup.GetPaginatedProvince)
	su.Get("/all/paginate/area/:area_uuid", sup.GetPaginatedArea)
	su.Post("/create", sup.CreateSup)
	su.Get("/get/:uuid", sup.GetSup)
	su.Put("/update/:uuid", sup.UpdateSup)
	su.Delete("/delete/:uuid", sup.DeleteSup)

	// Pos controller
	po := api.Group("/pos")
	po.Get("/all", pos.GetAllPoss)
	po.Get("/all/paginate", pos.GetPaginatedPos)
	po.Get("/all/paginate/province/:province_uuid", pos.GetPaginatedPosByProvinceUUID)
	po.Get("/all/paginate/area/:area_uuid", pos.GetPaginatedPosByAreaUUID)
	po.Get("/all/paginate/subarea/:subarea_uuid", pos.GetPaginatedPosBySubAreaUUID)
	po.Get("/all/paginate/:commune_uuid", pos.GetPaginatedPosByCommuneUUID)
	po.Get("/all/countries/:country_uuid", pos.GetAllPosByManager)
	po.Get("/all/provinces/:province_uuid", pos.GetAllPosByASM)
	po.Get("/all/areas/:area_uuid", pos.GetAllPosBySup)
	po.Get("/all/subareas/:sub_area_uuid", pos.GetAllPosByDR)
	po.Get("/all/cyclo/:cyclo_uuid", pos.GetAllPosByCyclo)
	po.Post("/create", pos.CreatePos)
	po.Get("/get/:uuid", pos.GetPos)
	po.Put("/update/:uuid", pos.UpdatePos)
	po.Delete("/delete/:uuid", pos.DeletePos)

	// Countries controller
	co := api.Group("/countries")
	co.Get("/all", country.GetAllCountry)
	co.Get("/all/paginate", country.GetPaginatedCountry)
	// co.Get("/all/dropdown", country.GetCountryDropdown)
	co.Get("/get/:uuid", country.GetCountry)
	co.Post("/create", country.CreateCountry)
	co.Put("/update/:uuid", country.UpdateCountry)
	co.Delete("/delete/:uuid", country.DeleteCountry)

	// DR Controller
	d := api.Group("/drs")
	d.Get("/all", dr.GetAllDr)
	d.Get("/all/paginate", dr.GetPaginatedDr)
	d.Get("/all/paginate/province/:province_uuid", dr.GetPaginatedDrByProvince)
	d.Get("/all/paginate/area/:area_uuid", dr.GetPaginatedDrByArea)
	d.Get("/all/paginate/subarea/:subarea_uuid", dr.GetPaginatedDrBySubArea)
	d.Get("/get/:uuid", dr.GetOneDr)
	d.Post("/create", dr.CreateDr)
	d.Put("/update/:uuid", dr.UpdateDr)
	d.Delete("/delete/:uuid", dr.DeleteDr)

	//Cyclo controller
	cy := api.Group("/cyclos")
	cy.Get("/all", cyclo.GetAllCyclo)
	cy.Get("/all/paginate", cyclo.GetPaginatedCyclo)
	cy.Get("/all/paginate/province/:province_uuid", cyclo.GetPaginatedCycloProvinceByID)
	cy.Get("/all/paginate/area/:area_uuid", cyclo.GetPaginatedCycloByAreaUUID)
	cy.Get("/all/paginate/subarea/:subarea_uuid", cyclo.GetPaginatedSubAreaByID)
	cy.Get("/all/paginate/commune/:user_uuid", cyclo.GetPaginatedCycloCommune)
	cy.Get("/get/:uuid", cyclo.GetOneCyclo)
	cy.Post("/create", cyclo.CreateCyclo)
	cy.Put("/update/:uuid", cyclo.UpdateCyclo)
	cy.Delete("/delete/:uuid", cyclo.DeleteCyclo)

	//routeplan controller
	rp := api.Group("/routeplans")
	rp.Get("/all", routeplan.GetRouteplan)
	rp.Get("/all/paginate", routeplan.GetPaginatedRouteplan)
	rp.Get("/all/paginate/province/:province_uuid", routeplan.GetPaginatedRouthplaByProvinceUUID)
	rp.Get("/all/paginate/area/:area_uuid", routeplan.GetPaginatedRouthplaByareaUUID)
	rp.Get("/all/paginate/subarea/:subarea_uuid", routeplan.GetPaginatedRouthplaBySubareaUUID)
	rp.Get("/all/paginate/:commune_uuid", routeplan.GetPaginatedRouteplaBycommuneUUID)
	rp.Get("/all/:id", routeplan.GetRouteplan)
	rp.Get("/get/:uuid", routeplan.GetRouteplan)
	rp.Post("/create", routeplan.CreateRouteplan)
	rp.Put("/update/:uuid", routeplan.UpdateRouteplan)
	rp.Delete("/delete/:uuid", routeplan.DeleteRouteplan)

	//routeplanitem controller
	rpi := api.Group("/routeplan-items")
	rpi.Get("/all", RoutePlanItem.GetPaginatedRoutePlanItem)
	rpi.Get("/all/paginate", RoutePlanItem.GetPaginatedRoutePlanItem)
	rpi.Get("/all/:id", RoutePlanItem.GetAllRoutePlanItem)
	rpi.Get("/get/:route_uuid", RoutePlanItem.GetOneByRouteUUID)
	rpi.Get("/get/:uuid", RoutePlanItem.GetAllRoutePlanItem)
	rpi.Post("/create", RoutePlanItem.CreateRoutePlanItem)
	rpi.Put("/update/:uuid", RoutePlanItem.UpdateRoutePlanItem)
	rpi.Delete("/delete/:uuid", RoutePlanItem.DeleteRoutePlanItem)

	// Brand controller
	br := api.Group("/brands")
	br.Get("/all", brand.GetAllBrands)
	br.Get("/all/provinces/:province_uuid", brand.GetAllBrandsByProvince)
	br.Get("/all/paginate", brand.GetPaginatedBrands)
	br.Get("/all/paginate/province/:province_uuid", brand.GetPaginatedBrandsByProvinceUUID)
	br.Get("/get/:uuid", brand.GetOneBrand)
	br.Post("/create", brand.CreateBrand)
	br.Put("/update/:uuid", brand.UpdateBrand)
	br.Delete("/delete/:uuid", brand.DeleteBrand)

	//POSEQUIPEMENT controller
	pe := api.Group("/posequipements")
	pe.Get("/all", posequiment.GetPaginatedPosEquipment)
	pe.Get("/all/paginate", posequiment.GetPaginatedPosEquipment)
	pe.Get("/all/:id", posequiment.GetPosEquipment)
	pe.Post("/create", posequiment.CreatePosEquipment)
	pe.Get("/get/:uuid", posequiment.GetAllPosEquipments)
	pe.Get("/get/:uuid", posequiment.GetPosEquipment)
	pe.Put("/update/:uuid", posequiment.UpdatePosEquipment)
	pe.Delete("/delete/:uuid", posequiment.DeletePosEquipment)

	// UserLogs controller
	log := api.Group("/users-logs")
	log.Get("/all", user_logs.GetUserLogs)
	log.Get("/all/paginate", user_logs.GetPaginatedUserLogs)
	log.Get("/all/paginate/:user_uuid", user_logs.GetUserLogByID)
	log.Get("/get/:uuid", user_logs.GetUserLog)
	log.Post("/create", user_logs.CreateUserLog)
	log.Put("/update/:uuid", user_logs.UpdateUserLog)
	log.Delete("/delete/:uuid", user_logs.DeleteUserLog)

	dash := api.Group("/dashboard")

	//ND Dashboard
	nd := dash.Group("/numeric-distribution")
	nd.Get("/table-view-province", dashboard.NdTableViewProvince)
	nd.Get("/table-view-area", dashboard.NdTableViewArea)
	nd.Get("/table-view-subarea", dashboard.NdTableViewSubArea)
	nd.Get("/table-view-commune", dashboard.NdTableViewCommune)
	nd.Get("/line-chart-by-month", dashboard.NdTotalByBrandByMonth)

	// SOS Dashboard
	sos := dash.Group("/share-of-stock")
	sos.Get("/table-view-province", dashboard.SosTableViewProvince)
	sos.Get("/table-view-area", dashboard.SosTableViewArea)
	sos.Get("/table-view-subarea", dashboard.SosTableViewSubArea)
	sos.Get("/table-view-commune", dashboard.SosTableViewCommune)
	sos.Get("/line-chart-by-month", dashboard.SosTotalByBrandByMonth)

	sum := dash.Group("/sammury")
	sum.Get("/dr-count", dashboard.DrCount)
	sum.Get("/pos-count", dashboard.POSCount)
	sum.Get("/province-count", dashboard.ProvinceCount)
	sum.Get("/area-count", dashboard.AreaCount)
	sum.Get("/sos-pie/:start_date/:end_date", dashboard.SOSPie)
	sum.Get("/tracking-visit-dr/:days/:start_date/:end_date", dashboard.TrackingVisitDRS)
	sum.Get("/summary-chart-bar/:start_date/:end_date", dashboard.SummaryChartBar)
	sum.Get("/better-dr/:start_date/:end_date", dashboard.BetterDR)
	sum.Get("/better-supervisor/:start_date/:end_date", dashboard.BetterSup)
	sum.Get("/status-equements/:start_date/:end_date", dashboard.StatusEquipement)
	sum.Get("/google-maps/:start_date/:end_date", dashboard.GoogleMaps)
	sum.Get("/price-sales/:start_date/:end_date", dashboard.PriceSale)

}

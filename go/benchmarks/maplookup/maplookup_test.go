package maplookup

import "testing"

type Permission int

const (
	None Permission = iota
	ConfigureSystem
	CreateUser
	UpdateUser
	DeleteUser
	CloseUser
	GrantWheelUserRole
	AssignUser
	UnassignUser
	SetUserPassword
	UpdateUserLocation
	CreateCustomer
	UpdateCustomer
	DeleteCustomer
	SetCustomerPassword
	CreateHub
	UpdateHub
	DeleteHub
	CreateBusiness
	UpdateBusiness
	CloseBusiness
	UpdateHubStock
	CreateMenu
	UpdateMenu
	DeleteMenu
	CreateMenuItem
	UpdateMenuItem
	DeleteMenuItem
	CreateIngredient
	UpdateIngredient
	DeleteIngredient
	CreateCopy
	UpdateCopy
	DeleteCopy
	CreateSalesOrder
	UpdateSalesOrder
	DeleteSalesOrder
	AssignSalesOrder
	ShipSalesOrder
	SetNextSalesOrder
	DeliverSalesOrder
	CancelSalesOrder
	RefundSalesOrder
	DeliverAnySalesOrder
	UpdateHubTicket
	GrantCredit
	CreateCreditDomain
	UpdateCreditDomain
	DeleteCreditDomain
	GrantInvite
	CreateDeliveratorRelease
	RequestDeliveratorLogs
	SubmitDeliveratorLogs
	CreateCampaign
	UpdateCampaign
	DeleteCampaign
	CreateCase
	UpdateCase
	CreateRecipe
	UpdateRecipe
	DeleteRecipe
	CreateVendorItem
	UpdateVendorItem
	DeleteVendorItem
	ModifyReward
	ModifyCustomerAnnouncement
	ModifyRegionMenuImage
	ModifyGuestChef
	ModifyMenuSection
	ActivateSelfDestructSequence
)

var perms = []Permission{
	UpdateUserLocation,
	UpdateCampaign,
	AssignUser,
	GrantWheelUserRole,
	ModifyMenuSection,
	ShipSalesOrder,
	CreateCopy,
	UpdateSalesOrder,
	CreateCreditDomain,
	UpdateHubStock,
	UpdateHubTicket,
	AssignSalesOrder,
	ModifyRegionMenuImage,
	DeleteMenu,
	CreateCase,
	UpdateCase,
	CreateCampaign,
	DeliverAnySalesOrder,
	CloseBusiness,
	RequestDeliveratorLogs,
	CreateDeliveratorRelease,
	CreateBusiness,
	DeleteHub,
	CreateCustomer,
	CreateMenu,
	UpdateRecipe,
	CancelSalesOrder,
	DeleteCopy,
	ConfigureSystem,
	DeleteCampaign,
	DeleteCustomer,
	DeliverSalesOrder,
	CreateHub,
}

type Set map[Permission]struct{}

func (s Set) add(p Permission) {
	s[p] = struct{}{}
}

func (s Set) contains(p Permission) bool {
	_, ok := s[p]
	return ok
}

type List []Permission

func (l List) add(p Permission) List {
	if l.contains(p) {
		return l
	}
	return append(l, p)
}

func (l List) contains(p Permission) bool {
	for _, q := range l {
		if p == q {
			return true
		}
	}
	return false
}

var (
	set  Set
	list List
)

func init() {
	set := make(Set)
	for _, p := range perms {
		set[p] = struct{}{}
	}
	list := make(List, len(perms))
	copy(list, perms)
}

/*
func BenchmarkSuccessfulSetLookupOne(b *testing.B) {
	p := UpdateUserLocation
	n := 0
	for i := 0; i < b.N; i++ {
		if set.contains(p) {
			n++
		}
	}
}

func BenchmarkUnsuccessfulSetLookupOne(b *testing.B) {
	p := UpdateVendorItem
	n := 0
	for i := 0; i < b.N; i++ {
		if set.contains(p) {
			n++
		}
	}
}

func BenchmarkSuccessfulSetLookupMany(b *testing.B) {
	ps := []Permission{
		ShipSalesOrder,
		CreateDeliveratorRelease,
		AssignUser,
		ModifyRegionMenuImage,
		CreateCopy,
		DeliverSalesOrder,
		DeleteHub,
		CreateMenu,
		CreateBusiness,
		DeleteMenu,
	}
	n := 0
	for i := 0; i < b.N; i++ {
		for _, p := range ps {
			if set.contains(p) {
				n++
			}
		}
	}
}

func BenchmarkSuccessfulListLookupFirstOne(b *testing.B) {
	p := UpdateUserLocation
	n := 0
	for i := 0; i < b.N; i++ {
		if list.contains(p) {
			n++
		}
	}
}

func BenchmarkSuccessfulListLookupMiddleOne(b *testing.B) {
	p := UpdateCase
	n := 0
	for i := 0; i < b.N; i++ {
		if list.contains(p) {
			n++
		}
	}
}

func BenchmarkSuccessfulListLookupLastOne(b *testing.B) {
	p := CreateHub
	n := 0
	for i := 0; i < b.N; i++ {
		if list.contains(p) {
			n++
		}
	}
}

func BenchmarkUnsuccessfulListLookupOne(b *testing.B) {
	p := UpdateVendorItem
	n := 0
	for i := 0; i < b.N; i++ {
		if list.contains(p) {
			n++
		}
	}
}

func BenchmarkSuccessfulListLookupMany(b *testing.B) {
	ps := []Permission{
		ShipSalesOrder,
		CreateDeliveratorRelease,
		AssignUser,
		ModifyRegionMenuImage,
		CreateCopy,
		DeliverSalesOrder,
		DeleteHub,
		CreateMenu,
		CreateBusiness,
		DeleteMenu,
	}
	n := 0
	for i := 0; i < b.N; i++ {
		for _, p := range ps {
			if list.contains(p) {
				n++
			}
		}
	}
}
*/

var roles = [][]Permission{
	{
		ModifyReward,
		CreateMenu,
		DeleteCopy,
		CreateRecipe,
		CreateIngredient,
		CreateDeliveratorRelease,
		UpdateHubTicket,
		UpdateHub,
	},
	{
		CloseBusiness,
		AssignUser,
		ModifyRegionMenuImage,
		DeleteIngredient,
		UpdateIngredient,
		CreateUser,
		CreateCase,
		CreateMenuItem,
	},
	{
		ShipSalesOrder,
		DeleteUser,
		UpdateCampaign,
		SetUserPassword,
		GrantWheelUserRole,
		DeleteCopy,
		AssignUser,
	},
	{
		DeleteSalesOrder,
		DeleteHub,
		UpdateUser,
		ModifyRegionMenuImage,
		DeleteCustomer,
		ModifyReward,
		CloseUser,
		UpdateCustomer,
	},
	{
		DeleteCreditDomain,
		UpdateHubTicket,
		CloseBusiness,
		CreateCampaign,
		SetUserPassword,
		UpdateMenu,
		CreateIngredient,
		ShipSalesOrder,
	},
}

func BenchmarkCreateSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := make(Set)
		for _, role := range roles {
			for _, p := range role {
				set.add(p)
			}
		}
	}
}

func BenchmarkCreateListNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := make(List, 0)
		for _, role := range roles {
			for _, p := range role {
				list.add(p)
			}
		}
	}
}

func BenchmarkCreateListPrealloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := 0
		for _, role := range roles {
			x += len(role)
		}
		list := make(List, 0, x)
		for _, role := range roles {
			for _, p := range role {
				list.add(p)
			}
		}
	}
}

func BenchmarkCreateListPreallocSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := 0
		for _, role := range roles {
			x += len(role)
		}
		flat := make([]Permission, 0, x)
		for _, role := range roles {
			flat = append(flat, role...)
		}
		list := make(List, 0, x)
		prev := None
		for i := 0; i < x; i++ {
			if flat[i] == prev {
				i++
			} else {
				prev = flat[i]
				list.add(prev)
			}
		}
	}
}

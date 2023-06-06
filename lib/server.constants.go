package lib

const (
	XApiKey      = "REPLACE_WITH_YOUR_KEY"
	XProductCode = "REPLACE_WITH_PRODUCT_CODE"
)

const GatewayHost = "https://gateway.snowpal.com/"

const (
	RouteRegistrationRegisterNewUserByEmail = "app/users/sign-up"
	RouteRegistrationSignInByEmail          = "app/users/sign-in"
	RouteRegistrationResetPassword          = "app/users/reset-password"
	RouteRegistrationActivateUser           = "app/user-verified/%s"
)

const (
	RouteUserGetUsers              = "users"
	RouteUserGetUserByUuid         = "users/uuid/%s"
	RouteUserDeactivateUserAccount = "users/deactivate-account"
	RouteUserDeleteUserAccount     = "users/delete-account"
)

const (
	RouteTeamsAddTeam    = "/teams"
	RouteTeamsGetTeams   = "/teams"
	RouteTeamsGetTeam    = "/teams/%s"
	RouteTeamsUpdateTeam = "/teams/%s"
	RouteTeamsDeleteTeam = "/teams/%s"
)

const (
	RouteMembersAddMemberToTeam      = "/teams/%s/members"
	RouteMembersGetTeamMembers       = "/teams/%s/members"
	RouteMembersDeleteMemberFromTeam = "/teams/%s/members"
)

const (
	RouteProvisioningCreateDatabaseIndexes = "/db/create-indexes"
)

const (
	RoutesMetaGetStatusFilters = "/status/filters"
	RoutesMetaGetRoles         = "/roles"
)

const (
	RouteProfilesAddProfile        = "/profiles"
	RouteProfilesGetProfileById    = "/profiles/by-id/%s"
	RouteProfilesGetProfileByEmail = "/profiles/by-email/%s"
	RouteProfilesGetMyProfile      = "/profiles"
	RouteProfilesUpdateProfile     = "/profiles"
)

const (
	RouteStatusesAddMyStatus          = "/teams/%s/statuses"
	RouteStatusesGetMyStatuses        = "/teams/%s/statuses"
	RouteStatusesGetAllMyStatuses     = "/statuses"
	RouteStatusesGetStatusById        = "/teams/%s/statuses/by-id/%s"
	RouteStatusesGetStatusByDate      = "/teams/%s/statuses/by-date/%s"
	RouteStatusesGetStatusByDateRange = "/teams/%s/statuses/by-date-range?startDate=%s&endDate=%s"
	RouteStatusesDeleteMyStatus       = "/teams/%s/statuses/%s"
)

const (
	RouteStatusesAddPftTickets       = "/teams/%s/statuses/%s/plan-for-today/tickets"
	RouteStatusesUpdatePftTicket     = "/teams/%s/statuses/%s/plan-for-today/tickets/%s"
	RouteStatusesUpdatePftOtherItems = "/teams/%s/statuses/%s/plan-for-today/other-items"
	RouteStatusesDeletePftTicket     = "/teams/%s/statuses/%s/plan-for-today/tickets/%s"
	RouteStatusesDeletePftOtherItems = "/teams/%s/statuses/%s/plan-for-today/other-items"
)

const (
	RouteStatusesAddSessionTickets      = "/teams/%s/statuses/%s/sessions/%s/tickets"
	RouteStatusesAddSessions            = "/teams/%s/statuses/%s/sessions"
	RouteStatusesAddSessionOtherItems   = "/teams/%s/statuses/%s/sessions/%s/other-items"
	RouteStatusesUpdateSessionTicket    = "/teams/%s/statuses/%s/sessions/%s/tickets/%s"
	RouteStatusesUpdateSessionOtherItem = "/teams/%s/statuses/%s/sessions/%s/other-items/%s"
	RouteStatusesDeleteSessionTicket    = "/teams/%s/statuses/%s/sessions/%s/tickets/%s"
	RouteStatusesDeleteSession          = "/teams/%s/statuses/%s/sessions/%s"
	RouteStatusesDeleteSessionOtherItem = "/teams/%s/statuses/%s/sessions/%s/other-items/%s"
)

const (
	RouteStatusesAddBlockedByTickets       = "/teams/%s/statuses/%s/blocked-by/tickets"
	RouteStatusesAddBlockedByMembers       = "/teams/%s/statuses/%s/blocked-by/members"
	RouteStatusesUpdateBlockedByTicket     = "/teams/%s/statuses/%s/blocked-by/tickets/%s"
	RouteStatusesUpdateBlockedByMember     = "/teams/%s/statuses/%s/blocked-by/members/%s"
	RouteStatusesUpdateBlockedByOtherItems = "/teams/%s/statuses/%s/blocked-by/other-items"
	RouteStatusesDeleteBlockedByTicket     = "/teams/%s/statuses/%s/blocked-by/tickets/%s"
	RouteStatusesDeleteBlockedByMember     = "/teams/%s/statuses/%s/blocked-by/members/%s"
	RouteStatusesDeleteBlockedByOtherItems = "/teams/%s/statuses/%s/blocked-by/other-items"
)

const (
	RouteStatusesAddStatusForMember            = "/teams/%s/members/%s/statuses"
	RouteStatusesGetTeamStatuses               = "/teams/%s/members/statuses"
	RouteStatusesGetStatusesForMember          = "/teams/%s/members/%s/statuses"
	RouteStatusesGetStatusByIdForMember        = "/teams/%s/statuses/by-id/%s"
	RouteStatusesGetStatusByDateForMember      = "/teams/%s/statuses/by-date/%s?memberId=%s"
	RouteStatusesGetStatusByDateRangeForMember = "/teams/%s/statuses/by-date-range?memberId=%s&startDate=%s&endDate=%s"
	RouteStatusesDeleteStatusForMember         = "/teams/%s/members/%s/statuses/%s"
)

const (
	RouteStatusesAddPftTicketsForMember       = "/teams/%s/members/%s/statuses/%s/plan-for-today/tickets"
	RouteStatusesUpdatePftTicketForMember     = "/teams/%s/members/%s/statuses/%s/plan-for-today/tickets/%s"
	RouteStatusesUpdatePftOtherItemsForMember = "/teams/%s/members/%s/statuses/%s/plan-for-today/other-items"
	RouteStatusesDeletePftTicketForMember     = "/teams/%s/members/%s/statuses/%s/plan-for-today/tickets/%s"
	RouteStatusesDeletePftOtherItemsForMember = "/teams/%s/members/%s/statuses/%s/plan-for-today/other-items"
)

const (
	RouteStatusesAddSessionTicketsForMember      = "/teams/%s/members/%s/statuses/%s/sessions/%s/tickets"
	RouteStatusesAddSessionsForMember            = "/teams/%s/members/%s/statuses/%s/sessions"
	RouteStatusesAddSessionOtherItemsForMember   = "/teams/%s/members/%s/statuses/%s/sessions/%s/other-items"
	RouteStatusesUpdateSessionTicketForMember    = "/teams/%s/members/%s/statuses/%s/sessions/%s/tickets/%s"
	RouteStatusesUpdateSessionOtherItemForMember = "/teams/%s/members/%s/statuses/%s/sessions/%s/other-items/%s"
	RouteStatusesDeleteSessionTicketForMember    = "/teams/%s/members/%s/statuses/%s/sessions/%s/tickets/%s"
	RouteStatusesDeleteSessionForMember          = "/teams/%s/members/%s/statuses/%s/sessions/%s"
	RouteStatusesDeleteSessionOtherItemForMember = "/teams/%s/members/%s/statuses/%s/sessions/%s/other-items/%s"
)

const (
	RouteStatusesAddBlockedByTicketsForMember       = "/teams/%s/members/%s/statuses/%s/blocked-by/tickets"
	RouteStatusesAddBlockedByMembersForMember       = "/teams/%s/members/%s/statuses/%s/blocked-by/members"
	RouteStatusesUpdateBlockedByTicketForMember     = "/teams/%s/members/%s/statuses/%s/blocked-by/tickets/%s"
	RouteStatusesUpdateBlockedByOtherItemsForMember = "/teams/%s/members/%s/statuses/%s/blocked-by/other-items"
	RouteStatusesUpdateBlockedByMemberForMember     = "/teams/%s/members/%s/statuses/%s/blocked-by/members/%s"
	RouteStatusesDeleteBlockedByTicketForMember     = "/teams/%s/members/%s/statuses/%s/blocked-by/tickets/%s"
	RouteStatusesDeleteBlockedByOtherItemsForMember = "/teams/%s/members/%s/statuses/%s/blocked-by/other-items"
	RouteStatusesDeleteBlockedByMemberForMember     = "/teams/%s/members/%s/statuses/%s/blocked-by/members/%s"
)

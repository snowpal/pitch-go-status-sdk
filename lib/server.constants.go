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
	RouteMembersDeleteMemberFromTeam = "/teams/%s/members/%s"
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
	RoutePullRequestsGetTeamPullRequests   = "/teams/%s/pull-requests"
	RoutePullRequestsGetMemberPullRequests = "/members/%s/pull-requests"
)

const (
	RouteStoryPointsGetTeamStoryPoints   = "/teams/%s/story-points"
	RouteStoryPointsGetMemberStoryPoints = "/members/%s/story-points"
)

const (
	RouteTicketsGetTeamTickets   = "/teams/%s/tickets"
	RouteTicketsGetMemberTickets = "/members/%s/story-points"
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
)

const (
	RouteStatusesAddSessionTickets       = "/teams/%s/statuses/%s/sessions/%s/tickets"
	RouteStatusesUpdateSessionTicket     = "/teams/%s/statuses/%s/sessions/%s/tickets/%s"
	RouteStatusesUpdateSessionOtherItems = "/teams/%s/statuses/%s/plan-for-today/other-items"
	RouteStatusesDeleteSessionTicket     = "/teams/%s/statuses/%s/sessions/%s/tickets/%s"
)

const (
	RouteStatusesAddBlockedByTickets  = "/teams/%s/statuses/%s/blocked-by/tickets"
	RouteStatusesUpdateBlockeByTicket = "/teams/%s/statuses/%s/blocked-by/tickets/%s"
	RouteStatusesDeleteBlockeByTicket = "/teams/%s/statuses/%s/blocked-by/tickets/%s"
)

const (
	RouteStatusesAddStatus = "/teams/%s/members/%s/statuses"
)

const (
	RouteStatusesAddPftTicketsForMember   = "/teams/%s/members/%s/statuses/%s/plan-for-today/tickets"
	RouteStatusesUpdatePftTicketForMember = "/teams/%s/members/%s/statuses/%s/plan-for-today/tickets/%s"
	RouteStatusesDeletePftTicketForMember = "/teams/%s/members/%s/statuses/%s/plan-for-today/tickets/%s"
)

const (
	RouteStatusesAddSessionTicketsForMember   = "/teams/%s/members/%s/statuses/%s/sessions/%s/tickets"
	RouteStatusesUpdateSessionTicketForMember = "/teams/%s/members/%s/statuses/%s/sessions/%s/tickets/%s"
	RouteStatusesDeleteSessionTicketForMember = "/teams/%s/members/%s/statuses/%s/sessions/%s/tickets/%s"
)

const (
	RouteStatusesAddBlockedByTicketsForMember  = "/teams/%s/members/%s/statuses/%s/blocked-by/tickets"
	RouteStatusesUpdateBlockeByTicketForMember = "/teams/%s/members/%s/statuses/%s/blocked-by/tickets/%s"
	RouteStatusesDeleteBlockeByTicketForMember = "/teams/%s/members/%s/statuses/%s/blocked-by/tickets/%s"
)

const (
	RouteStatusesGetMemberStatuses  = "/members/%s/statuses"
	RouteStatusesGetMemberStatus    = "/members/%s/statuses/%s"
	RouteStatusesGetAllTeamStatuses = "/teams/statuses"
	RouteStatusesUpdateMemberStatus = "/members/%s/statuses/%s"
	RouteStatusesUpdateMyStatus     = "/statuses/%s"
	RouteStatusesDeleteMyStatuses   = "/statuses"
	RouteStatusesDeleteTeamStatuses = "/teams/%s/statuses"
)

const (
	RouteSearchSrchStatuses = "/search/statuses"
	RouteSearchSrchTeams    = "/search/teams"
	RouteSearchSrchMembers  = "/search/members"
)

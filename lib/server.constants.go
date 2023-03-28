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
	RouteMembersGetTeamMember        = "/teams/%s/members/%s"
	RouteMembersGetBlockedMembers    = "/teams/%s/members/blocked"
	RouteMembersGetAllBlockedMembers = "/teams/members/blocked"
	RouteMembersUpdateTeamMember     = "/teams/%s/members/%s"
	RouteMembersDeleteTeamMember     = "/teams/%s/members/%s"
)

const (
	RouteRolesAddRole    = "/roles"
	RouteRolesGetRoles   = "/roles"
	RouteRolesUpdateRole = "/roles/%s"
	RouteRolesDeleteRole = "/roles/%s"
)

const (
	RouteCommentsAddComment    = "/status/%s/comments"
	RouteCommentsGetComments   = "/status/%s/comments"
	RouteCommentsGetComment    = "/status/%s/comments/%s"
	RouteCommentsUpdateComment = "/status/%s/comments/%s"
	RouteCommentsDeleteComment = "/status/%s/comments/%s"
)

const (
	RoutesMetaGetStatusFilters = "/status/%s/filters"
)

const (
	RouteProfilesGetProfile    = "/profiles/%s"
	RouteProfilesUpdateProfile = "/profiles/%s"
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
	RouteStatusesAddMemberStatus      = "/members/%s/statuses"
	RouteStatusesGetMemberStatuses    = "/members/%s/statuses"
	RouteStatusesGetMemberStatus      = "/members/%s/statuses/%s"
	RouteStatusesGetTeamStatuses      = "/teams/%s/statuses"
	RouteStatusesGetAllTeamStatuses   = "/teams/statuses"
	RouteStatusesUpdateMemberStatus   = "/members/%s/statuses/%s"
	RouteStatusesDeleteMemberStatuses = "/members/%s/statuses"
	RouteStatusesDeleteTeamStatuses   = "/teams/%s/statuses"
)

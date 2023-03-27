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
	RouteCommentsAddComment    = "/status/%s/comments"
	RouteCommentsGetComments   = "/status/%s/comments"
	RouteCommentsGetComment    = "/status/%s/comments/%s"
	RouteCommentsUpdateComment = "/status/%s/comments/%s"
	RouteCommentsDeleteComment = "/status/%s/comments/%s"
)

const (
	RoutesMetaGetStatusFilters = "/status/%s/comments"
)

const (
	RouteUserGetUsers              = "users"
	RouteUserGetUserByUuid         = "users/uuid/%s"
	RouteUserDeactivateUserAccount = "users/deactivate-account"
)

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Yamashou/gqlgenc/clientv2"
	cato "github.com/routebyintuition/cato-go-sdk"
	cato_models "github.com/routebyintuition/cato-go-sdk/models"
)

func main() {
	token := os.Getenv("CATO_API_KEY")
	accountId := os.Getenv("CATO_ACCOUNT_ID")
	url := "https://api.catonetworks.com/api/v1/graphql2"

	if token == "" {
		fmt.Println("no token provided")
		os.Exit(1)
	}

	if accountId == "" {
		fmt.Println("no account id provided")
		os.Exit(1)
	}

	catoClient := &cato.Client{
		Client: clientv2.NewClient(http.DefaultClient, url, nil,
			func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {
				req.Header.Set("x-api-key", token)

				return next(ctx, req, gqlInfo, res)
			}),
	}

	ctx := context.Background()

	position := cato_models.PolicyRulePositionEnum("LAST_IN_POLICY")
	hostRefInput := []*cato_models.HostRef{}
	siteRefInput := []*cato_models.SiteRef{}
	iprange := []*cato_models.IPAddressRangeInput{}
	globalIpRange := []*cato_models.GlobalIPRangeRef{}
	networkInterfaceRefInput := []*cato_models.NetworkInterfaceRef{}
	siteNetworkSubnetRefInput := []*cato_models.SiteNetworkSubnetRef{}
	floatingSubnetRefInput := []*cato_models.FloatingSubnetRef{}
	userRefInput := []*cato_models.UserRef{}
	usersGroupRefInput := []*cato_models.UsersGroupRef{}
	groupRefInput := []*cato_models.GroupRef{}
	systemGroupRefInput := []*cato_models.SystemGroupRef{}

	connectionOrigin := cato_models.ConnectionOriginEnum("ANY")
	actionEnum := cato_models.InternetFirewallActionEnum("ALLOW")

	domainList := []string{"slashdot.org"}
	fqdnList := []string{"www.slashdot.org"}

	inputRule := cato_models.InternetFirewallAddRuleInput{
		At: &cato_models.PolicyRulePositionInput{
			Position: &position,
		},
		Rule: &cato_models.InternetFirewallAddRuleDataInput{
			Enabled:     false,
			Name:        "TestRule101",
			Description: "TestRule101",
			Source: &cato_models.InternetFirewallSourceInput{
				IP:                []string{},
				Host:              hostRefInput,
				Site:              siteRefInput,
				Subnet:            []string{},
				IPRange:           iprange,
				GlobalIPRange:     globalIpRange,
				NetworkInterface:  networkInterfaceRefInput,
				SiteNetworkSubnet: siteNetworkSubnetRefInput,
				FloatingSubnet:    floatingSubnetRefInput,
				User:              userRefInput,
				UsersGroup:        usersGroupRefInput,
				Group:             groupRefInput,
				SystemGroup:       systemGroupRefInput,
			},
			ConnectionOrigin: connectionOrigin,
			Country:          []*cato_models.CountryRef{},
			Device:           []*cato_models.DeviceProfileRef{},
			DeviceOs:         []cato_models.OperatingSystem{},
			Destination: &cato_models.InternetFirewallDestinationInput{
				Application:            []*cato_models.ApplicationRef{},
				CustomApp:              []*cato_models.CustomApplicationRef{},
				AppCategory:            []*cato_models.ApplicationCategoryRef{},
				CustomCategory:         []*cato_models.CustomCategoryRef{},
				SanctionedAppsCategory: []*cato_models.SanctionedAppsCategoryRef{},
				Country:                []*cato_models.CountryRef{},
				Domain:                 domainList,
				Fqdn:                   fqdnList,
				IP:                     []string{},
				Subnet:                 []string{},
				IPRange:                iprange,
				GlobalIPRange:          globalIpRange,
				RemoteAsn:              []string{},
			},
			Service: &cato_models.InternetFirewallServiceTypeInput{},
			Action:  actionEnum,
			Schedule: &cato_models.PolicyScheduleInput{
				ActiveOn: "ALWAYS",
			},
			Tracking: &cato_models.PolicyTrackingInput{
				Event: &cato_models.PolicyRuleTrackingEventInput{
					Enabled: true,
				},
				Alert: &cato_models.PolicyRuleTrackingAlertInput{
					Enabled:           false,
					Frequency:         "DAILY",
					SubscriptionGroup: []*cato_models.SubscriptionGroupRef{},
					MailingList:       []*cato_models.SubscriptionMailingListRef{},
					Webhook:           []*cato_models.SubscriptionWebhookRef{},
				},
			},
		},
	}

	policyChange, err := catoClient.PolicyInternetFirewallAddRule(ctx, inputRule, accountId)

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	fmt.Println("policyChange: ", policyChange)

}

package main

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"gopkg.in/gomail.v2"
)

func mainwws() {
	toEmails := []string{"mayur.wadekar@pb.com"}
	ccEmails := []string{}  //"wadekarmh@gmail.com"}
	bccEmails := []string{} //"wadekarmh@gmail.com"}
	attachments := []string{}
	fromEmail := "shipping.au@pb.com"

	// new session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		fmt.Println("Error occurred while creating aws session", err)
		return
	}

	// create recipients
	var recipients []*string
	for _, r := range toEmails {
		recipient := r
		recipients = append(recipients, &recipient)
	}

	// Create an SES session.
	svc := ses.New(sess)

	// create raw message
	msg := gomail.NewMessage()
	msg.SetAddressHeader("From", fromEmail, "Australia")
	msg.SetHeader("To", toEmails...)
	msg.SetHeader("Subject", "Sample subject")
	msg.SetBody("text/html", `
	<!DOCTYPE html\r\n PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional //EN\"
    \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd\">\r\n<html xmlns=\"http://www.w3.org/1999/xhtml\"
    style=\"box-sizing: border-box; font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size: 16px;
    margin: 0px; padding: 0px\">\r\n\r\n <head>\r\n<meta charset=\"UTF-8\" />\r\n
    <meta http-equiv=\"Content-Type\" name=\"viewport\" content=\"text/html; charset=utf-8; initial-scale=1\" />\r\n
    <title>Welcome to SendPro<sup style=\"box-sizing: border-box; font-family: 'Helvetica Neue'
        , Helvetica, Arial, sans-serif; font-size: .6em; margin: 0px; padding: 0px\">™</sup> Online</title>\r\n\r\n
</head>\r\n\r\n
<body style=\"-webkit-font-smoothing: antialiased; -webkit-text-size-adjust: none; background: #f0f0f0; box-sizing:
    border-box; font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size: 16px; margin: 0px; padding:
    0px; width: 100% !important\" bgcolor=\"#f0f0f0\">\r\n <style type=\"text/css\">
        \r\n @media (max-width: 768px) {
            \r\n .container-fluid {
                \r\n padding-left: 15px;
                \r\n padding-right: 15px;
                \r\n
            }

            \r\n
        }

        \r\n\r\n body {
            \r\n -webkit-font-smoothing: antialiased;
            \r\n -webkit-text-size-adjust: none;
            \r\n width: 100% !important;
            \r\n background-color: #f0f0f0;
            \r\n
        }

        \r\n\r\n img {
            \r\n max-width: 100%;
            \r\n
        }

        \r\n\r\n @media only screen and (max-device-width: 480px) {
            \r\n table[class=\"container\"] {
                \r\n width: 90% !important;
                \r\n
            }

            \r\n\r\n .header-logo-cell {
                \r\n padding: 0 !important;
                \r\n min-width: 100% !important;
                \r\n height: auto !important;
                \r\n display: block !important;
                \r\n text-align: center !important;
                \r\n clear: both !important;
                \r\n
            }

            \r\n\r\n .header-logo {
                \r\n margin: 0 !important;
                \r\n padding: 0 !important;
                \r\n text-align: center !important;
                \r\n width: 175px !important;
                \r\n height: 44px !important;
                \r\n
            }

            \r\n\r\n .header-product-name-cell {
                \r\n padding: 0 !important;
                \r\n min-width: 100% !important;
                \r\n height: auto !important;
                \r\n display: block !important;
                \r\n text-align: center !important;
                \r\n clear: both !important;
                \r\n
            }

            \r\n\r\n .header-product-name {
                \r\n display: block !important;
                \r\n margin: 20px 0 0 0 !important;
                \r\n padding: 0 !important;
                \r\n text-align: center !important;
                \r\n
            }

            \r\n\r\n .body h1 {
                \r\n font-size: 24px !important;
                \r\n
            }

            \r\n
        }

        \r\n .pdf {
            color: blue !important;
            text-decoration: underline !important;
            \r\n
        }

        \r\n
    </style>\r\n <div class=\"wrapper\" style=\"background: #f0f0f0; box-sizing: border-box;
        font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size: 16px; margin: 0px; overflow: visible;
        padding: 0px; width: 100%\">\r\n\r\n <table class=\"container\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\"
            width=\"580\" align=\"center\" style=\"-webkit-font-smoothing: antialiased; -webkit-text-size-adjust: none;
            background: #f0f0f0; box-sizing: border-box; clear: both !important; display: block !important;
            font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size: 16px; margin: 0 auto; padding: 0px;
            width: 580px\" bgcolor=\"#f0f0f0\">\r\n <tr style=\"box-sizing: border-box; font-family: 'Helvetica Neue' ,
                Helvetica, Arial, sans-serif; font-size: 16px; margin: 0px; padding: 0px\">\r\n <th
                    class=\"header-spacer\" style=\"box-sizing: border-box; font-family: 'Helvetica Neue' , Helvetica,
                    Arial, sans-serif; font-size: 16px; margin: 0px; padding: 0\">\r\n \r\n </th>\r\n </tr>\r\n\r\n <tr
                style=\"box-sizing: border-box; font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size:
                16px; margin: 0px; padding: 0px\">\r\n <td style=\"box-sizing: border-box; font-family: 'Helvetica Neue'
                    , Helvetica, Arial, sans-serif; font-size: 16px; margin: 0px; padding: 0px\">\r\n <table
                        border=\"0\" cellspacing=\"0\" cellpadding=\"0\" width=\"100%\" style=\"box-sizing: border-box;
                        font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size: 16px; margin: 0px;
                        padding: 0px\">\r\n <tr style=\"box-sizing: border-box; font-family: 'Helvetica Neue' ,
                            Helvetica, Arial, sans-serif; font-size: 16px; margin: 0px; padding: 0px\">\r\n <th
                                class=\"header-logo-cell\" height=\"30\" align=\"left\" valign=\"middle\"
                                style=\"box-sizing: border-box; font-family: 'Helvetica Neue' , Helvetica, Arial,
                                sans-serif; font-size: 16px; margin: 0px; padding: 0px 0px 0px 30px; width: 50%\">\r\n
                                <img class=\"header-logo\" alt=\"Pitney Bowes\" width=\"120\" height=\"24\"
                                    src=\"https://cdn.designsystem.pitneycloud.com/email_test_images/logo.gif\"
                                    style=\"box-sizing: border-box; font-family: 'Helvetica Neue' , Helvetica, Arial,
                                    sans-serif; font-size: 16px; margin: 0px; max-width: 100%; padding: 0px\" />\r\n
                            </th>\r\n <th class=\"header-product-name-cell\" height=\"30\" align=\"right\"
                                valign=\"middle\" style=\"box-sizing: border-box; font-family: 'Helvetica Neue' ,
                                Helvetica, Arial, sans-serif; font-size: 16px; margin: 0px; padding: 0px 30px 0px 0px;
                                width: 50%\">\r\n \r\n </th>\r\n </tr>\r\n </table>\r\n
                </td>\r\n </tr>\r\n\r\n <tr style=\"box-sizing: border-box; font-family: 'Helvetica Neue' , Helvetica,
                Arial, sans-serif; font-size: 16px; margin: 0px; padding: 0px\">\r\n <th class=\"header-spacer\"
                    style=\"box-sizing: border-box; font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif;
                    font-size: 16px; margin: 0px; padding: 0\">\r\n \r\n </th>\r\n </tr>\r\n\r\n <tr
                class=\"header-gradient-row\" style=\"box-sizing: border-box; font-family: 'Helvetica Neue' , Helvetica,
                Arial, sans-serif; font-size: 0; line-height: 0; margin: 0px; padding: 0px\">\r\n <th
                    class=\"header-gradient\" height=\"16\" width=\"100%\" style=\"background: #3e53a4; box-sizing:
                    border-box; font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size: 0; line-height:
                    0; margin: 0px; padding: 0px\" bgcolor=\"#3e53a4\">\r\n <img alt=\"Pitney Bowes\" height=\"16\"
                        src=\"http://www.pbdesignsystem.com/email_test_images/header_gradient.gif\" style=\"box-sizing:
                        border-box; font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size: 16px;
                        margin: 0px; width: 100%; padding: 0px\" />\r\n\r\n </th>\r\n </tr>\r\n\r\n <tr
                style=\"box-sizing: border-box; font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size:
                16px; margin: 0px; padding: 0px\">\r\n <td class=\"body\" valign=\"middle\" width=\"100%\"
                    style=\"background: #fff; box-sizing: border-box; font-family: 'Helvetica Neue' , Helvetica, Arial,
                    sans-serif; font-size: 16px; margin: 0px; padding: 50px 30px\" bgcolor=\"#fff\">\r\n <h1
                        style=\"box-sizing: border-box; color: #cf0989; font-family: 'Helvetica Neue' , Helvetica,
                        Arial, sans-serif; font-size: 32px; font-weight: normal; line-height: 1em; margin: 0px 0px 30px;
                        padding: 0px\">Welcome to SendPro<sup style=\"box-sizing: border-box; font-family: 'Helvetica Neue'
                        , Helvetica, Arial, sans-serif; font-size: .6em; margin: 0px; padding: 0px\">™</sup> Online!</h1>\r\n\r\n <p style=\"box-sizing: border-box; clear:
                        both; color: #717171; font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size:
                        16px; line-height: 26px; margin: 10px 0px 0px; padding: 0px\">Hi {{FirstName}}, </p>\r\n\r\n 
                        <!-- <p
                        style=\"box-sizing: border-box; clear: both; color: #717171; font-family: 'Helvetica Neue' ,
                        Helvetica, Arial, sans-serif; font-size: 16px; line-height: 26px; margin: 10px 0px 0px; padding:
                        0px\">Thank you for signing up to SendPro<sup style=\"box-sizing: border-box; font-family: 'Helvetica Neue'
                        , Helvetica, Arial, sans-serif; font-size: .6em; margin: 0px; padding: 0px\">™</sup> Online – your new, easy-to-use<br> online shipping
                        platform.\r\n </p>\r\n\r\n <br /> -->
                    <p style=\"box-sizing: border-box; clear: both; color: #717171; font-family: 'Helvetica Neue' ,
                        Helvetica, Arial, sans-serif; font-size: 16px; line-height: 26px; margin: 10px 0px 0px; padding:
                        0px\"> All you need to do is sign in using your existing email {{EmailAddress}}<br> and your Pitney Bowes account password. </p>\r\n <p class=\"button-row\"
                        style=\"box-sizing: border-box; clear: both; color: #717171; font-family: 'Helvetica Neue' ,
                        Helvetica, Arial, sans-serif; font-size: 16px; line-height: 26px; margin: 30px 0px; padding:
                        0px\">\r\n
                        <!-- MUST ADJUST WIDTH AND TEXT IN BUTTON HERE -->\r\n
                        <!--[if mso]> <v:roundrect xmlns:v=\"urn:schemas-microsoft-com:vml\" xmlns:w=\"urn:schemas-microsoft-com:office:word\" href=\"{{WelcomeLink}}\" style=\"height:45px;v-text-anchor:middle;width:300px;\" arcsize=\"50%\" stroke=\"f\" fillcolor=\"#3E53A4\"> <w:anchorlock/> <center> <![endif]-->\r\n
                        <!-- Change token and link below for SPOG -->\r\n\t\t\t<a href=\"{{WelcomeLink}}\"
                            style=\"-webkit-text-size-adjust: none; background: #3E53A4; border-radius: 4px; box-sizing:
                            border-box; color: #ffffff; display: inline-block; font-family: sans-serif; font-size: 16px;
                            line-height: 45px; margin: 0px; padding: 0px; text-align: center; text-decoration: none;
                            width: 300px\">Login</a>\r\n
                        <!--[if mso]> </center> </v:roundrect> <![endif]-->\r\n
                    </p>\r\n\r\n \r\n <p style=\"box-sizing: border-box; clear: both; color: #717171;
                        font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size: 16px; line-height:
                        26px; margin: 10px 0px 0px; padding: 0px\"> The SendPro<sup style=\"box-sizing: border-box; font-family: 'Helvetica Neue'
                        , Helvetica, Arial, sans-serif; font-size: .6em; margin: 0px; padding: 0px\">™</sup> Online PAYG plan lets you compare shipping prices, print<br> shipping labels and track packages. Get started immediately by using our<br>
                         discount shipping rates with Aramex and CouriersPlease. </p>\r\n\r\n
                    <br>\r\n
                </p>\r\n\r\n \r\n <p style=\"box-sizing: border-box; clear: both; color: #717171;
                font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size: 16px; line-height:
                26px; margin: 10px 0px 0px; padding: 0px\"> You can also integrate your online stores for simple and secure sending<br> and automated order management. </p>\r\n\r\n
            <br>\r\n

                    <b><span style=\"box-sizing: border-box; clear: both; color: #717171; font-family: 'Helvetica Neue'
                            , Helvetica, Arial, sans-serif; font-size: 16px; line-height: 26px; margin: 10px 0px 0px;
                            padding: 0px\">Here’s how to get shipping:</span></b> <br>\r\n
                    <br>\r\n
                    <b><span style=\"box-sizing: border-box; clear: both; color: #717171; font-family: 'Helvetica Neue'
                            , Helvetica, Arial, sans-serif; font-size: 16px; line-height: 26px; margin: 10px 0px 0px;
                            padding: 0px\">Step 1: Finalise your account set-up</span></b><br>
                    <p style=\"box-sizing: border-box; clear: both; color: #717171; font-family: 'Helvetica Neue' ,
                        Helvetica, Arial, sans-serif; font-size: 16px; line-height: 26px; margin: 10px 0px 0px; padding:
                        0px\">Before you can ship a package, you’ll need to add a payment method.<br> You can add your card details by clicking on the Cards/Payment tab<br>
                         under your name.  \r\n </p>\r\n\r\n
                    <br>\r\n \r\n\r\n 
                    <b><span style=\"box-sizing: border-box; clear: both; color: #717171;
                            font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size: 16px; line-height:
                            26px; margin: 10px 0px 0px; padding: 0px\">Step 2: Create a shipping label</span></b><br>
                    <ul style=\"box-sizing: border-box; color: #717171; font-family: 'Helvetica Neue' , Helvetica,
                        Arial, sans-serif; font-size: 16px; line-height: 26px; margin: 10px 0px 0px 20px; padding:
                        0px\">\r\n<li style=\"box-sizing: border-box; font-family: 'Helvetica Neue' , Helvetica, Arial,
                            sans-serif; font-size: 16px; margin: 0px; padding: 0px\">\r\n Click on SendPro<sup style=\"box-sizing: border-box; font-family: 'Helvetica Neue'
                            , Helvetica, Arial, sans-serif; font-size: .6em; margin: 0px; padding: 0px\">™</sup> Online in the
                            left top corner to load the home<br> screen where you can choose your package size. \r\n
                        </li>\r\n <li style=\"box-sizing: border-box; font-family: 'Helvetica Neue' , Helvetica, Arial,
                            sans-serif; font-size: 16px; margin: 0px; padding: 0px\">\r\n Select the appropriate size or
                            enter custom dimensions \r\n </li>\r\n <li style=\"box-sizing: border-box;
                            font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size: 16px; margin: 0px;
                            padding: 0px\">\r\n Then just fill out the details for your shipment </li>\r\n <li
                            style=\"box-sizing: border-box; font-family: 'Helvetica Neue' , Helvetica, Arial,
                            sans-serif; font-size: 16px; margin: 0px; padding: 0px\">\r\n Print label </li>\r\n <li
                            style=\"box-sizing: border-box; font-family: 'Helvetica Neue' , Helvetica, Arial,
                            sans-serif; font-size: 16px; margin: 0px; padding: 0px\">\r\n Get ready for pick-up of your
                            package </li>\r\n </ul>\r\n <br>\r\n \r\n\r\n <b><span style=\"box-sizing: border-box;
                            clear: both; color: #717171; font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif;
                            font-size: 16px; line-height: 26px; margin: 10px 0px 0px; padding: 0px\">Step 3: Connect
                            your online store</span></b><br>
                    <p style=\"box-sizing: border-box; clear: both; color: #717171; font-family: 'Helvetica Neue' ,
                        Helvetica, Arial, sans-serif; font-size: 16px; line-height: 26px; margin: 10px 0px 0px; padding:
                        0px\">If you have an online store, you can easily connect that to SendPro<sup style=\"box-sizing: border-box; font-family: 'Helvetica Neue'
                        , Helvetica, Arial, sans-serif; font-size: .6em; margin: 0px; padding: 0px\">™</sup><br> Online to automate
                        your order management. Check out our <a class=\"pdf\"
                            href=\"https://www.pitneybowes.com/content/dam/pitneybowes/australia/en/campaigns/ecommerce-stores-integration-guide.pdf\"
                            target=\"_blank\" style=\"border: 0; border-bottom-color: #717171; border-box; color:
                            #717171; font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size: 16px;
                            margin: 0px; padding: 10px; \">‘eCommerce<br> stores integration guide’</a> to help connect
                        your eBay, Shopify and<br> WooCommerce stores. \r\n </p>\r\n\r\n <br>\r\n \r\n\r\n
                        <b><span
                            style=\"box-sizing: border-box; clear: both; color: #717171; font-family: 'Helvetica Neue' ,
                            Helvetica, Arial, sans-serif; font-size: 16px; line-height: 26px; margin: 10px 0px 0px;
                            padding: 0px\">With SendPro<sup style=\"box-sizing: border-box; font-family: 'Helvetica Neue'
                            , Helvetica, Arial, sans-serif; font-size: .6em; margin: 0px; padding: 0px\">™</sup> Online you can: </span></b><br>
                    <ul style=\"box-sizing: border-box; color: #717171; font-family: 'Helvetica Neue' , Helvetica,
                        Arial, sans-serif; font-size: 16px; line-height: 26px; margin: 10px 0px 0px 20px; padding:
                        0px\">\r\n<li style=\"box-sizing: border-box; font-family: 'Helvetica Neue' , Helvetica, Arial,
                            sans-serif; font-size: 16px; margin: 0px; padding: 0px\">\r\n Easily print shipping labels
                            \r\n </li>\r\n <li style=\"box-sizing: border-box; font-family: 'Helvetica Neue' ,
                            Helvetica, Arial, sans-serif; font-size: 16px; margin: 0px; padding: 0px\">\r\n Compare
                            shipping prices from different carriers \r\n </li>\r\n <li style=\"box-sizing: border-box;
                            font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size: 16px; margin: 0px;
                            padding: 0px\">\r\n Integrate your online stores (Shopify, eBay and WooCommerce) </li>\r\n
                        <li style=\"box-sizing: border-box; font-family: 'Helvetica Neue' , Helvetica, Arial,
                            sans-serif; font-size: 16px; margin: 0px; padding: 0px\">\r\n Send and track shipments from
                            departure to arrival </li>\r\n <li style=\"box-sizing: border-box;
                            font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size: 16px; margin: 0px;
                            padding: 0px\">\r\n Access discounted carrier rates with no minimum volume </li>\r\n </ul>
                    \r\n <br>\r\n\r\n \r\n <p style=\"box-sizing: border-box; clear: both; color: #717171;
                        font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size: 16px; line-height:
                        26px; margin: 10px 0px 0px; padding: 0px\">If you need help to start shipping, contact us via
                        chat on<br> spo.pitneybowes.com, via <a class=\"pdf\" href=\"mailto:SendPro_Online_AU@pb.com\"
                            target=\"_blank\" style=\"border: 0; border-bottom-color: #717171; border-box; color:
                            #717171; font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size: 16px;
                            margin: 0px; padding: 10px; \">email</a> or call us on 13 23 63. </p> \r\n\r\n
                </td>\r\n </tr>\r\n\r\n
            <!-- EMAIL FOOTER -->\r\n <tr style=\"box-sizing: border-box; font-family: 'Helvetica Neue' , Helvetica,
                Arial, sans-serif; font-size: 16px; margin: 0px; padding: 0px\">\r\n \r\n <td class=\"footer\"
                    valign=\"middle\" width=\"100%\" align=\"left\" style=\"box-sizing: border-box;
                    font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size: 16px; margin: 0px; padding:
                    0px 30px\">\r\n <p style=\"box-sizing: border-box; color: #717171; font-family: 'Helvetica Neue' ,
                        Helvetica, Arial, sans-serif; font-size: 10px; margin: 0px 0px 5px; padding: 0\">\r\n <a
                            class=\"footer-logo\" href=\"http://www.pitneybowes.com/au\" target=\"_blank\"
                            style=\"border: 0; border-bottom-color: #717171; border-bottom-style: solid; box-sizing:
                            border-box; color: #717171; font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif;
                            font-size: 10px; margin: 0px; padding: 10px; text-decoration: none\">\r\n </a>\r\n </p>\r\n
                    <p style=\"box-sizing: border-box; color: #717171; font-family: 'Helvetica Neue' , Helvetica, Arial,
                        sans-serif; font-size: 10px; margin: 0px 0px 5px; padding: 0\">This email is an auto-generated
                        message from an unmonitored email account. To get in touch with Pitney Bowes,<br> email <a
                            class=\"pdf\" href=\"mailto:SendPro_Online_AU@pb.com\" target=\"_blank\" style=\"box-sizing:
                            border-box; color: #717171; font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif;
                            font-size: 10px; margin: 0px 0px 5px; padding: 0\">SendPro_Online_AU@pb.com </a>\r\n
                    <p style=\"box-sizing: border-box; color: #717171; font-family: 'Helvetica Neue' , Helvetica, Arial,
                        sans-serif; font-size: 10px; margin: 0px 0px 5px; padding: 0\">This message was distributed by
                        Pitney Bowes, Level 1, 68 Waterloo Road, Macquarie Park NSW 2113 Australia. <br> Pitney Bowes,
                        the Corporate Logo and SendPro<sup style=\"box-sizing: border-box; font-family: 'Helvetica Neue'
                            , Helvetica, Arial, sans-serif; font-size: .6em; margin: 0px; padding: 0px\">™</sup> Online
                        are trademarks of Pitney Bowes Inc. or a subsidiary. All other<br> marks are the intellectual
                        property of their respective owners. Use of your email address is governed by the Pitney Bowes
                        <a href=\"https://www.pitneybowes.com/au/privacy-statement.html\" style=\"border-bottom-color:
                            #717171; border-bottom-style: solid; border-bottom-width: 1px; box-sizing: border-box;
                            color: #717171; font-family: 'Helvetica Neue' , Helvetica, Arial, sans-serif; font-size:
                            10px; margin: 0px; padding: 0px; text-decoration: none\">Privacy Policy</a>.</p>\r\n <p
                        style=\"box-sizing: border-box; color: #717171; font-family: 'Helvetica Neue' , Helvetica,
                        Arial, sans-serif; font-size: 10px; margin: 0px 0px 5px; padding: 0\">©2021 Pitney Bowes Inc.
                        All rights reserved.</p>\r\n
                </td>\r\n </tr>\r\n\r\n
        </table>\r\n\r\n </div>\r\n</body>\r\n\r\n

</html>
	`)

	// cc mails mentioned
	if len(ccEmails) != 0 {
		// Need to add cc mail IDs also in recipient list
		for _, r := range ccEmails {
			recipient := r
			recipients = append(recipients, &recipient)
		}
		msg.SetHeader("cc", ccEmails...)
	}

	// As per documentation need to add CC and BCC emails with To field of email. Due to only single destination list.

	// bcc mails mentioned
	if len(bccEmails) != 0 {
		// Need to add bcc mail IDs also in recipient list
		for _, r := range bccEmails {
			recipient := r
			recipients = append(recipients, &recipient)
		}
		msg.SetHeader("bcc", bccEmails...)
	}

	// If attachments exists
	if len(attachments) != 0 {
		for _, f := range attachments {
			msg.Attach(f)
		}
	}

	// create a new buffer to add raw data
	var emailRaw bytes.Buffer
	msg.WriteTo(&emailRaw)

	// create new raw message
	message := ses.RawMessage{Data: emailRaw.Bytes()}

	input := &ses.SendRawEmailInput{Source: &fromEmail, Destinations: recipients, RawMessage: &message}

	// send raw email
	_, err = svc.SendRawEmail(input)
	if err != nil {
		fmt.Println("Error ", err)
	}
}

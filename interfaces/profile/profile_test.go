package profile

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"pinterest/application"
	"pinterest/application/mock_application"
	"pinterest/domain/entity"
	"pinterest/interfaces/auth"
	"testing"
	"time"

	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	"pinterest/interfaces/middleware"
)

// profileInputStruct stores information which will be parsed into request
type profileInputStruct struct {
	url          string
	urlForRouter string
	method       string
	headers      map[string][]string
	postBody     []byte // JSON
	profileFunc  func(w http.ResponseWriter, r *http.Request)
	middleware   func(next http.HandlerFunc, cookieApp application.CookieAppInterface) http.HandlerFunc
}

// toHTTPRequest transforms profileInputStruct to http.Request, adding global cookies
func (input *profileInputStruct) toHTTPRequest(cookies []*http.Cookie) *http.Request {
	reqURL, _ := url.Parse("http://localhost:8080" + input.url) // Scheme (http://) is required for URL parsing
	reqBody := bytes.NewBuffer(input.postBody)
	request := &http.Request{
		Method:        input.method,
		URL:           reqURL,
		Header:        input.headers,
		ContentLength: int64(reqBody.Len()),
		Body:          ioutil.NopCloser(reqBody),
	}

	if (len(cookies) > 0) && (request.Header == nil) {
		request.Header = make(http.Header)
	}

	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}

	return request
}

// profileOutputStruct stores information parsed from response
type profileOutputStruct struct {
	responseCode int
	headers      map[string][]string
	postBody     []byte // JSON
}

// fillFromResponse transforms http.Response to profileOutputStruct
func (output *profileOutputStruct) fillFromResponse(response *http.Response) error {
	output.responseCode = response.StatusCode
	output.headers = response.Header
	if len(output.headers) == 0 {
		output.headers = nil
	}
	var err error
	output.postBody, err = ioutil.ReadAll(response.Body)
	if len(output.postBody) == 0 {
		output.postBody = nil
	}
	return err
}

var testProfileInfo ProfileInfo
var testAuthInfo auth.AuthInfo

// These tests have to run in that order!!!
var profileTestSuccess = []struct {
	in   profileInputStruct
	out  profileOutputStruct
	name string
}{
	{
		profileInputStruct{
			"/auth/signup",
			"/auth/signup",
			"POST",
			nil,
			[]byte(`{"username": "TestUsername",` +
				`"password": "thisisapassword",` +
				`"email": "test@example.com",` +
				`"firstName": "TestFirstName",` +
				`"lastName": "TestLastName",` +
				`"avatarLink": "avatars/1"}`,
			),
			testAuthInfo.HandleCreateUser,
			middleware.NoAuthMid,
		},

		profileOutputStruct{
			201,
			nil,
			nil,
		},
		"Testing profile creation",
	},
	{
		profileInputStruct{
			"/profile",
			"/profile",
			"GET",
			nil,
			nil,
			testProfileInfo.HandleGetProfile,
			middleware.AuthMid, // If user is not logged in, they can't access their profile
		},

		profileOutputStruct{
			200,
			nil,
			[]byte(`{"ID":0,` +
				`"username":"TestUsername",` +
				`"email":"test@example.com",` +
				`"firstName":"TestFirstName",` +
				`"lastName":"TestLastName",` +
				`"avatarLink":"avatars/1",` +
				`"following":0,` +
				`"followers":0}`,
			),
		},
		"Testing profile output",
	},
	{
		profileInputStruct{
			"/profile/password",
			"/profile/password",
			"PUT",
			nil,
			[]byte(`{"password":"New Password"}`),
			testProfileInfo.HandleChangePassword,
			middleware.AuthMid,
		},

		profileOutputStruct{
			204,
			nil,
			nil,
		},
		"Testing password change", // I don't know right now how to easily check if password changed
	},
	{
		profileInputStruct{
			"/profile/TestUsername",
			"/profile/{username}",
			"GET",
			nil,
			nil,
			testProfileInfo.HandleGetProfile,
			nil,
		},

		profileOutputStruct{
			200,
			nil,
			[]byte(`{"ID":0,` +
				`"username":"TestUsername",` +
				`"email":"test@example.com",` +
				`"firstName":"TestFirstName",` +
				`"lastName":"TestLastName",` +
				`"avatarLink":"avatars/1",` +
				`"following":0,` +
				`"followers":0}`,
			),
		},
		"Testing profile output using profile name",
	},
	{
		profileInputStruct{
			"/profile/edit",
			"/profile/edit",
			"PUT",
			nil,
			[]byte(`{"username": "new_User_Name",` +
				`"firstName": "new First name",` +
				`"lastName": "new Last Name",` +
				`"email": "new@example.com",` +
				`"avatarLink": "avatars/2"}`,
			),
			testProfileInfo.HandleEditProfile,
			middleware.AuthMid,
		},

		profileOutputStruct{
			204,
			nil,
			nil,
		},
		"Testing profile edit",
	},
	{
		profileInputStruct{
			"/profile/0",
			"/profile/{id:[0-9]+}",
			"GET",
			nil,
			nil,
			testProfileInfo.HandleGetProfile,
			nil,
		},

		profileOutputStruct{
			200,
			nil,
			[]byte(`{"ID":0,` +
				`"username":"new_User_Name",` +
				`"email":"new@example.com",` +
				`"firstName":"new First name",` +
				`"lastName":"new Last Name",` +
				`"avatarLink":"avatars/2",` +
				`"following":0,` +
				`"followers":0}`,
			),
		},
		"Testing profile output using profile id",
	},
	{
		profileInputStruct{
			"/profile/avatar",
			"/profile/avatar",
			"PUT",
			map[string][]string{
				"Content-Type": {"multipart/form-data; boundary=---------------------------9051914041544843365972754266"},
			},
			[]byte(`-----------------------------9051914041544843365972754266` + "\n" +
				`Content-Disposition: form-data; name="avatarImage"; filename="a.txt"` + "\n" +
				`Content-Type: image/jpeg` + "\n" +
				"\n" +
				`randomImage` + "\n" +
				"\n" +
				`-----------------------------9051914041544843365972754266--` + "\n"),
			testProfileInfo.HandlePostAvatar,
			middleware.AuthMid,
		},

		profileOutputStruct{
			204,
			nil,
			nil,
		},
		"Testing avatar change",
	},
	{
		profileInputStruct{
			"/profile/delete",
			"/profile/delete",
			"DELETE",
			nil,
			nil,
			testProfileInfo.HandleDeleteProfile,
			middleware.AuthMid,
		},

		profileOutputStruct{
			204,
			nil,
			nil,
		},
		"Testing profile deletion",
	},
}

var successCookies []*http.Cookie

func TestProfileSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserApp := mock_application.NewMockUserAppInterface(mockCtrl)
	cookieApp := application.NewCookieApp(40, 10*time.Hour)
	mockS3App := mock_application.NewMockS3AppInterface(mockCtrl)
	mockNotificationApp := mock_application.NewMockNotificationAppInterface(mockCtrl)

	// TODO: maybe replace this with JSON parsing?
	expectedUser := entity.User{
		UserID:    0,
		Username:  "TestUsername",
		Password:  "thisisapassword",
		FirstName: "TestFirstName",
		LastName:  "TestLastName",
		Email:     "test@example.com",
		Avatar:    "avatars/1",
		Salt:      "",
	}

	mockUserApp.EXPECT().GetUserByUsername(gomock.Any()).Return(nil, fmt.Errorf("No user found with such username")).Times(1) // Handler will request user info
	mockUserApp.EXPECT().CreateUser(gomock.Any()).Return(expectedUser.UserID, nil).Times(1)
	mockNotificationApp.EXPECT().ChangeToken(expectedUser.UserID, gomock.Any()).Return(nil).Times(1) // Adding notification token during user creation

	mockUserApp.EXPECT().GetUser(expectedUser.UserID).Return(&expectedUser, nil).Times(1) // Normal user output using cookie's userID

	mockUserApp.EXPECT().GetUser(expectedUser.UserID).Return(&expectedUser, nil).Times(1) // Before changing password, handler requests user data
	expectedUser.Password = "New Password"
	mockUserApp.EXPECT().SaveUser(gomock.Any()).Return(nil).Times(1) // Password changing

	mockUserApp.EXPECT().GetUserByUsername(expectedUser.Username).Return(&expectedUser, nil).Times(1) // Normal user output using username

	mockUserApp.EXPECT().GetUser(expectedUser.UserID).Return(&expectedUser, nil).Times(1) // Before profile editing, handler reuqests user data
	expectedUserEdited := entity.User{
		UserID:    0,
		Username:  "new_User_Name",
		Password:  "New Password",
		FirstName: "new First name",
		LastName:  "new Last Name",
		Email:     "new@example.com",
		Avatar:    "avatars/2",
		Salt:      "",
	}
	mockUserApp.EXPECT().SaveUser(gomock.Any()).Return(nil).Times(1) // Profile editing

	mockUserApp.EXPECT().GetUser(expectedUserEdited.UserID).Return(&expectedUserEdited, nil).Times(1) // Normal user output using userID

	mockUserApp.EXPECT().UpdateAvatar(expectedUser.UserID, gomock.Any()).Return(nil).Times(1)

	mockUserApp.EXPECT().DeleteUser(expectedUserEdited.UserID).Return(nil).Times(1)

	testAuthInfo = *auth.NewAuthInfo(
		mockUserApp,
		cookieApp,
		nil,
		nil,
		mockNotificationApp,
	)

	testProfileInfo = ProfileInfo{
		userApp:         mockUserApp,
		cookieApp:       cookieApp,
		s3App:           mockS3App,
		notificationApp: mockNotificationApp,
	}
	for _, tt := range profileTestSuccess {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			req := tt.in.toHTTPRequest(successCookies)

			rw := httptest.NewRecorder() // not ResponseWriter because we need to read response
			m := mux.NewRouter()
			funcToHandle := tt.in.profileFunc
			if tt.in.middleware != nil { // We don't always need middleware
				funcToHandle = tt.in.middleware(funcToHandle, cookieApp)
			}
			m.HandleFunc(tt.in.urlForRouter, funcToHandle).Methods(tt.in.method)
			m.ServeHTTP(rw, req)
			resp := rw.Result()

			// if server returned cookies, we use them
			if len(resp.Cookies()) > 0 {
				successCookies = resp.Cookies()
			}

			var result profileOutputStruct
			result.fillFromResponse(resp)

			require.Equal(t, tt.out.responseCode, result.responseCode,
				fmt.Sprintf("Expected: %d as response code\nbut got:  %d",
					tt.out.responseCode, result.responseCode))
			for key, val := range tt.out.headers {
				resultVal, ok := result.headers[key]
				require.True(t, !ok,
					fmt.Sprintf("Expected header %s is not found:\nExpected: %v\nbut got: %v", key, tt.out.headers, result.headers))
				require.Equal(t, val, resultVal,
					fmt.Sprintf("Expected value of header %s: %v is different from actual value: %v", key, val, resultVal))
			}
			require.Equal(t, tt.out.postBody, result.postBody,
				fmt.Sprintf("Expected: %v as response body\nbut got:  %v",
					string(tt.out.postBody), string(result.postBody)))
		})
	}
}

// These tests have to run in that order!!!
var profileTestFailure = []struct {
	in   profileInputStruct
	out  profileOutputStruct
	name string
}{
	{
		profileInputStruct{
			"/auth/signup",
			"/auth/signup",
			"POST",
			nil,
			[]byte(`{"username": "TestUsername",` +
				`"password": "thisisapassword",` +
				`"email": "test@example.com",` +
				`"firstName": "TestFirstName",` +
				`"lastName": "TestLastName",` +
				`"avatarLink": "avatars/1"}`,
			),
			testAuthInfo.HandleCreateUser,
			middleware.NoAuthMid,
		},

		profileOutputStruct{
			201,
			nil,
			nil,
		},
		"Testing profile creation",
	},
	{
		profileInputStruct{
			"/profile/password",
			"/profile/password",
			"PUT",
			nil,
			[]byte(`{"password":"SomeInvalidPassword"}}`), //JSON itself is incorrect
			testProfileInfo.HandleChangePassword,
			middleware.AuthMid,
		},

		profileOutputStruct{
			400,
			nil,
			nil,
		},
		"Testing password change if JSON is incorrect",
	},
	{
		profileInputStruct{
			"/profile/password",
			"/profile/password",
			"PUT",
			nil,
			[]byte(`{"password":"invalid"}`), // Password should be 8 characters or more long
			testProfileInfo.HandleChangePassword,
			middleware.AuthMid,
		},

		profileOutputStruct{
			400,
			nil,
			nil,
		},
		"Testing password change if password is invalid",
	},
	{
		profileInputStruct{
			"/profile/edit",
			"/profile/edit",
			"PUT",
			nil,
			[]byte(`{{"username": "new_User_Name",` + // Notice 2 opening brackets
				`"firstName": "new First name",` +
				`"lastName": "new Last Name",` +
				`"email": "new@example.com",` +
				`"avatarLink": "avatars/2"}`,
			),
			testProfileInfo.HandleEditProfile,
			middleware.AuthMid,
		},

		profileOutputStruct{
			400,
			nil,
			nil,
		},
		"Testing profile edit if JSON is incorrect",
	},
	{
		profileInputStruct{
			"/profile/edit",
			"/profile/edit",
			"PUT",
			nil,
			[]byte(`{"username": "new_User_Name",` +
				`"firstName": "new First name",` +
				`"lastName": "new Last Name",` +
				`"email": "new@example.com",` +
				`"avatarLink": "avatars/2"}`,
			),
			testProfileInfo.HandleEditProfile,
			middleware.AuthMid,
		},

		profileOutputStruct{
			409,
			nil,
			nil,
		},
		"Testing profile edit if username is not unique", // Username not being unique is entirely mock's prerogative
	},
}

var failureCookies []*http.Cookie

func TestProfileFailure(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserApp := mock_application.NewMockUserAppInterface(mockCtrl)
	cookieApp := application.NewCookieApp(40, 10*time.Hour)
	mockS3App := mock_application.NewMockS3AppInterface(mockCtrl)
	mockNotificationApp := mock_application.NewMockNotificationAppInterface(mockCtrl)

	expectedUser := entity.User{
		UserID:    0,
		Username:  "TestUsername",
		Password:  "thisisapassword",
		FirstName: "TestFirstName",
		LastName:  "TestLastName",
		Email:     "test@example.com",
		Avatar:    entity.AvatarDefaultPath,
		Salt:      "",
	}

	mockUserApp.EXPECT().GetUserByUsername(gomock.Any()).Return(nil, fmt.Errorf("No user found with such username")).Times(1) // Handler will request user info
	mockUserApp.EXPECT().CreateUser(gomock.Any()).Return(expectedUser.UserID, nil).Times(1)
	mockNotificationApp.EXPECT().ChangeToken(expectedUser.UserID, gomock.Any()).Return(nil).Times(1) // Adding notification token during user creation

	// During password change, if anything is wrong with JSON input, handler does not interact with database, hence no mocks

	mockUserApp.EXPECT().GetUser(expectedUser.UserID).Return(&expectedUser, nil).Times(1)
	mockUserApp.EXPECT().SaveUser(gomock.Any()).Return(fmt.Errorf("Username or email is already taken")).Times(1)

	testAuthInfo = *auth.NewAuthInfo(
		mockUserApp,
		cookieApp,
		nil, // We don't need S3 bucket in these tests
		nil, // We don't really care about boards in these tests
		mockNotificationApp,
	)
	testProfileInfo = ProfileInfo{
		userApp:         mockUserApp,
		cookieApp:       cookieApp,
		s3App:           mockS3App,
		notificationApp: mockNotificationApp,
	}

	for _, tt := range profileTestFailure {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			req := tt.in.toHTTPRequest(failureCookies)

			rw := httptest.NewRecorder() // not ResponseWriter because we need to read response
			m := mux.NewRouter()
			funcToHandle := tt.in.profileFunc
			if tt.in.middleware != nil { // We don't always need middleware
				funcToHandle = tt.in.middleware(funcToHandle, cookieApp)
			}
			m.HandleFunc(tt.in.urlForRouter, funcToHandle).Methods(tt.in.method)
			m.ServeHTTP(rw, req)
			resp := rw.Result()

			// if server returned cookies, we use them
			if len(resp.Cookies()) > 0 {
				failureCookies = resp.Cookies()
			}

			var result profileOutputStruct
			result.fillFromResponse(resp)

			require.Equal(t, tt.out.responseCode, result.responseCode,
				fmt.Sprintf("Expected: %d as response code\nbut got:  %d",
					tt.out.responseCode, result.responseCode))
			for key, val := range tt.out.headers {
				resultVal, ok := result.headers[key]
				require.True(t, !ok,
					fmt.Sprintf("Expected header %s is not found:\nExpected: %v\nbut got: %v", key, tt.out.headers, result.headers))
				require.Equal(t, val, resultVal,
					fmt.Sprintf("Expected value of header %s: %v is different from actual value: %v", key, val, resultVal))
			}
			require.Equal(t, tt.out.postBody, result.postBody,
				fmt.Sprintf("Expected: %v as response body\nbut got:  %v",
					string(tt.out.postBody), string(result.postBody)))
		})
	}
}

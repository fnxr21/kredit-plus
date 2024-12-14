package std

func GetMessage(code string) string {
	if msg, e := Message[code]; e {
		return msg
	}
	return "unknown_error"
}

var Message = map[string]string{
	// (penghuni/subpenghuni),
	"id_not_found":            "not-found", //
	"incorrect_password":      "incorrect-password",
	"admin_account_not_found": "admin-account-not-found",
}

// GetMessage returns the response for a given message code, or a default if not found

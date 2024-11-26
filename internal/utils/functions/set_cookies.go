package functions

import (
	"banners_oto/config"
	"banners_oto/internal/usecase/session"
	"banners_oto/internal/utils/alias"
	cnst "banners_oto/internal/utils/constants"
	"net/http"
	"time"

	"github.com/satori/uuid"
)

const timeExpDur = 14 * 24 * time.Hour

// SetCookie - !!it is necessary to take into account the time zone!!
func SetCookie(w http.ResponseWriter, r *http.Request, sessionClient session.Usecase, email string, cfg *config.ProjectConfiguration) (http.ResponseWriter, error) {
	sessionId := uuid.NewV4().String()
	err := sessionClient.SetValue(r.Context(), alias.SessionKey(sessionId), alias.SessionValue(email), int32(cfg.Redis.DatabaseSession))
	if err != nil {
		return w, err
	}
	expiration := time.Now().Add(timeExpDur)
	cookie := http.Cookie{
		Name:     cnst.SessionCookieName,
		Value:    sessionId,
		Expires:  expiration,
		HttpOnly: false,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)

	csrfToken, err := GenerateCsrfToken(cfg.Server.CsrfSecretKey, alias.SessionKey(sessionId))
	if err != nil {
		return w, err
	}
	err = sessionClient.SetValue(r.Context(), alias.SessionKey(csrfToken), alias.SessionValue(email), int32(cfg.Redis.DatabaseCsrf))
	if err != nil {
		return w, err
	}
	csrfCookie := http.Cookie{
		Name:     cnst.CsrfCookieName,
		Value:    csrfToken,
		Expires:  expiration,
		HttpOnly: false,
		Path:     "/",
	}
	http.SetCookie(w, &csrfCookie)
	return w, nil
}

package functions

import (
	"banners_oto/config"
	"banners_oto/internal/delivery/metrics"
	"banners_oto/internal/usecase/session"
	"banners_oto/internal/utils/alias"
	cnst "banners_oto/internal/utils/constants"
	"banners_oto/internal/utils/myerrors"
	"errors"
	"net/http"
)

// DeleteCookiesFromDB - method deletes session_id and csrf_token from Redis dbs
func DeleteCookiesFromDB(r *http.Request, ucSession session.Usecase, cfg *config.Redis, m *metrics.Metrics) error {
	sessionCookie, err := r.Cookie(cnst.SessionCookieName)
	if err != nil {
		return err
	}

	err = ucSession.DeleteKey(r.Context(), alias.SessionKey(sessionCookie.Value), int32(cfg.DatabaseSession))
	if err != nil {
		return err
	}

	csrfCookie, err := r.Cookie(cnst.CsrfCookieName)
	if err != nil {
		return err
	}
	err = ucSession.DeleteKey(r.Context(), alias.SessionKey(csrfCookie.Value), int32(cfg.DatabaseCsrf))
	if err != nil {
		return err
	}
	return nil
}

// CookieExpired - method set cookie expired
func CookieExpired(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, error) {
	sessionCookie := &http.Cookie{
		Name:     cnst.SessionCookieName,
		Value:    "",
		MaxAge:   -1,
		HttpOnly: false,
		Path:     "/",
	}
	http.SetCookie(w, sessionCookie)

	csrfCookie := &http.Cookie{
		Name:     cnst.CsrfCookieName,
		Value:    "",
		MaxAge:   -1,
		HttpOnly: false,
		Path:     "/",
	}
	http.SetCookie(w, csrfCookie)

	return w, nil
}

func FlashCookie(r *http.Request, w http.ResponseWriter, ucSession session.Usecase, cfg *config.Redis, m *metrics.Metrics) (http.ResponseWriter, error) {
	err := DeleteCookiesFromDB(r, ucSession, cfg, m)
	if err != nil {
		if !(errors.Is(err, myerrors.RedisNoData) || errors.Is(err, http.ErrNoCookie)) {
			return w, err
		}
	}

	w, err = CookieExpired(w, r)
	if err != nil && !errors.Is(err, http.ErrNoCookie) {
		return w, err
	}
	return w, nil
}

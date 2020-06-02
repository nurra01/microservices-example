/**
 * sets access token in local storage
 * @param {string} token - access token
 */
export function setAccessToken(token) {
    localStorage.setItem('access_token', token)
}
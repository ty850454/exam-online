const baseurl = 'http://127.0.0.1:8080'

export function post(url, data) {
    return fetch(baseurl + url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjE5LCJ1c2VybmFtZSI6IjEyMyJ9.MsjsL9DZzhPr5UUGf9iIyjKyEwByNJUWpX7xct24pUc'
        },
        body: JSON.stringify(data)
    }).then(response => {
        if (response.status >= 200 && response.status < 400) {
            return response.json()
        }
        return Promise.reject(response.json())
    })
}
export function get(url) {
    return fetch(baseurl + url, {
        method: 'GET',
        headers: {
            'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjE5LCJ1c2VybmFtZSI6IjEyMyJ9.MsjsL9DZzhPr5UUGf9iIyjKyEwByNJUWpX7xct24pUc'
        },
    }).then(response => {
        if (response.status >= 200 && response.status < 400) {
            return response.json()
        }
        return Promise.reject(response)
    })
}

export function patch(url) {
    return fetch(baseurl + url, {
        method: 'PATCH',
        headers: {
            'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjE5LCJ1c2VybmFtZSI6IjEyMyJ9.MsjsL9DZzhPr5UUGf9iIyjKyEwByNJUWpX7xct24pUc'
        },
    }).then(response => {
        if (response.status >= 200 && response.status < 400) {
            return response.json()
        }
        return Promise.reject(response)
    })
}

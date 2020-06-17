import request from '@/utils/simple-request'

const urlPrefix = process.env.VUE_APP_BASE_ALERTM_API

export function login(data) {
  return request({
    url: `${urlPrefix}/user/login`,
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: `${urlPrefix}/user/info`,
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    url: `${urlPrefix}/user/logout`,
    method: 'post'
  })
}

export function getAllUser() {
  return request({
    url: `${urlPrefix}/users`,
    method: 'get'
  })
}

export function updateUser(email, data) {
  return request({
    url: `${urlPrefix}/user/${email}`,
    method: 'put',
    data
  })
}

export function deleteUser(email) {
  return request({
    url: `${urlPrefix}/user/${email}`,
    method: 'delete'
  })
}

export function getAllUsergroup() {
  return request({
    url: `${urlPrefix}/ugroup`,
    method: 'get'
  })
}

export function createUsergroup(data) {
  return request({
    url: `${urlPrefix}/ugroup`,
    method: 'post',
    data
  })
}

export function updateUsergroup(group, data) {
  return request({
    url: `${urlPrefix}/ugroup/${group}`,
    method: 'put',
    data
  })
}

export function deleteUsergroup(group) {
  return request({
    url: `${urlPrefix}/ugroup/${group}`,
    method: 'delete'
  })
}

export function getAlertgroup() {
  return request({
    url: `${urlPrefix}/item`,
    method: 'get'
  })
}

export function createAlertgroup(data) {
  return request({
    url: `${urlPrefix}/item`,
    method: 'post',
    data
  })
}

export function updateAlertgroup(group, data) {
  return request({
    url: `${urlPrefix}/item/${group}`,
    method: 'put',
    data
  })
}

export function deleteAlertgroup(group) {
  return request({
    url: `${urlPrefix}/item/${group}`,
    method: 'delete'
  })
}

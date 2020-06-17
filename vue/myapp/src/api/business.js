import request from '@/utils/simple-request'

export function searchLive() {
  return request({
    url: '/live/zhenai',
    method: 'get'
  })
}

export function searchDevice() {
  return request({
    url: '/device',
    method: 'get'
  })
}

export function searchCDN() {
  return request({
    url: '/cdn',
    method: 'get'
  })
}

export function searchNginxDomain() {
  return request({
    url: '/nginx',
    method: 'get'
  })
}

export function searchNginxXX() {
  return request({
    url: '/nginx/day',
    method: 'get'
  })
}

export function searchCeSu() {
  return request({
    url: '/cesu',
    method: 'get'
  })
}

export function searchNginxPV() {
  return request({
    url: '/nginx/day',
    method: 'get'
  })
}

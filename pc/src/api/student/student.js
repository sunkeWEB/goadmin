import request from '@/utils/request'

export function getStudent(data) {
  return request({
    url: '/api/v1/student/getPage',
    method: 'get',
    params: data
  })
}

export function addStudent(data) {
  data.sex = Number(data.sex)
  return request({
    url: '/api/v1/student/add',
    method: 'post',
    data: data
  })
}

export function updateStudent(data) {
  data.sex = Number(data.sex)
  return request({
    url: '/api/v1/student/update',
    method: 'put',
    data: data
  })
}

export function deleteStudent(data) {
  return request({
    url: '/api/v1/student/deleteStudent',
    method: 'put',
    data: data
  })
}


import fetch from '@/utils/fetch'

export function addProject(data) {
  return fetch({
    url: 'project/create',
    method: 'POST',
    data
  })
}

export function updateProject(data) {
  return fetch({
    url: 'project/update',
    method: 'POST',
    data
  })
}

export function deleteProject(data) {
  return fetch({
    url: 'project/delete',
    method: 'POST',
    data
  })
}

export function projectList() {
  return fetch({
    url: 'project/list',
    method: 'get'
  })
}

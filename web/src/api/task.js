import fetch from '@/utils/fetch'

export function addTask(data) {
  return fetch({
    url: 'task/create',
    method: 'POST',
    data
  })
}

export function updateTask(data) {
  return fetch({
    url: 'task/update',
    method: 'POST',
    data
  })
}

export function deleteTask(data) {
  return fetch({
    url: 'task/delete',
    method: 'POST',
    data
  })
}

export function taskList() {
  return fetch({
    url: 'task/list',
    method: 'get'
  })
}

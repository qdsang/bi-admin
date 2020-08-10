import fetch from '@/utils/fetch'

export function addDashboard(data) {
  return fetch({
    url: '/dashboard/create',
    method: 'POST',
    data
  })
}

export function updateDashboard(data) {
  return fetch({
    url: '/dashboard/update',
    method: 'POST',
    data
  })
}

export function getdDashboardById(id) {
  return fetch({
    url: `/dashboard/get/${id}`
  })
}

export function deleteDashboard(data) {
  return fetch({
    url: `/dashboard/delete`,
    method: 'post',
    data
  })
}

export function dashboardList(projectId) {
  return fetch({
    url: `/dashboard/list?project_id=${projectId}`
  }).then((resp) => {
    resp.data.dashboards = resp.data.dashboards || resp.data.list
    resp.data.dashboards.forEach((item, index) => {
      if (typeof item.content === 'string') {
        try {
          item.content = JSON.parse(item.content)
        } catch (ex) {
          // console.log('dashboardList err json', ex)
        }
      }
    })
    return resp
  })
}

export function addChartToDB(data) {
  return fetch({
    url: '/chartboard/map',
    method: 'POST',
    data
  })
}

export function chartByDashboard(id) {
  return fetch({
    url: `/chartboardmap/chartbydashboard?dashboard_id=${id}`
  }).then((resp) => {
    return resp
  })
}

export function dbByChart(id) {
  return fetch({
    url: `/chartboardmap/boardbychart?chart_id=${id}`
  })
}

export function unMapChartDb(data) {
  return fetch({
    url: '/chartboard/unmap',
    method: 'POST',
    data
  })
}

export function dbOrder(data) {
  return fetch({
    url: '/dashboard/order',
    method: 'POST',
    data
  })
}

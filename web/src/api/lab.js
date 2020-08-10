import fetch from '@/utils/fetch'

export function addLab(data) {
  return fetch({
    url: '/lab/create',
    method: 'POST',
    data
  })
}

export function updateLab(data) {
  return fetch({
    url: '/lab/update',
    method: 'POST',
    data
  })
}

export function getdLabById(id) {
  return fetch({
    url: `/lab/${id}`
  })
}

export function deleteLab(data) {
  return fetch({
    url: `/lab/delete`,
    method: 'post',
    data
  })
}

export function labList() {
  return fetch({
    url: `/lab/list`
  }).then((resp) => {
    resp.data.labs = resp.data.labs || resp.data.list
    resp.data.labs.forEach((item, index) => {
      if (typeof item.content === 'string') {
        try {
          item.content = JSON.parse(item.content)
        } catch (ex) {
          // console.log('labList err json', ex)
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

export function chartByLab(id) {
  return fetch({
    url: `/chartboardmap/chartbylab?lab_id=${id}`
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
    url: '/lab/order',
    method: 'POST',
    data
  })
}

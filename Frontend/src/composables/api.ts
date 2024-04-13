type httpMethod =  'GET' | 'POST' | 'PUT' | 'DELETE'

export async function useFetch(fetchRoute: string, data: object, method:httpMethod = 'GET') {
  const BASE_URL = 'http://localhost:8080';
  const url = `${BASE_URL}/${fetchRoute}`;

  const requestOptions: RequestInit = {
    method: method
  }

  if (method !== 'GET') requestOptions.body = JSON.stringify(data);

  const res = await fetch(url, requestOptions);
  if (!res.ok) {
    const errorResponse = await res.json();
    throw new Error(errorResponse.error);
  }
  return res.json();
}

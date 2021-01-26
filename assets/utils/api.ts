export const fetchApi = async (url: string, params: any = {}) => {
  params = {
    header: {
      'Content-Type': 'application/json',
      Accept: 'application/json',
    },
    body: JSON.stringify(params.body),
    ...params,
  };

  const res = await fetch(url, params);
  return res.json();
};

export default fetchApi;

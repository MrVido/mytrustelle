// utils/api.ts

const BASE_URL = 'https://yourapi.domain.com';

export const fetchData = async (endpoint: string, params: Record<string, string> = {}) => {
  try {
    const url = new URL(endpoint, BASE_URL);
    Object.keys(params).forEach((key) => url.searchParams.append(key, params[key]));

    const response = await fetch(url.toString(), {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        // Include other headers as needed
      },
    });

    if (!response.ok) {
      throw new Error(`Error: ${response.status}`);
    }

    return await response.json();
  } catch (error) {
    console.error('Error fetching data:', error);
    throw error;
  }
};

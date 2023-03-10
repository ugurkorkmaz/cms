import type { PageLoad } from './$types';
 
export const load = (async () => {
  return {
    meta: {
      title: `Login Page`,
    }
  };
}) satisfies PageLoad;
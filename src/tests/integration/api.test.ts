import { Client } from '../../api/client';

describe('API', () => {
  it('should return all items when perPage is zero', async () => {
    const client = new Client();
    const items = await client.list('items', { perPage: 0 });
    expect(items.length).toBeGreaterThan(0);
  });

  it('should return a 400 Bad Request error for negative perPage', async () => {
    const client = new Client();
    await expect(client.list('items', { perPage: -1 })).rejects.toThrow('Bad Request');
  });

  it('should return a 400 Bad Request error for non-integer perPage', async () => {
    const client = new Client();
    await expect(client.list('items', { perPage: 1.5 })).rejects.toThrow('Bad Request');
  });
});

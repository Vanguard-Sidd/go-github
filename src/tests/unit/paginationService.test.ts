import { paginate } from '../../services/paginationService';

describe('paginationService', () => {
  it('should return all items when perPage is zero', () => {
    const items = [1, 2, 3, 4, 5];
    const options = { perPage: 0 };
    const result = paginate(items, options);
    expect(result).toEqual(items);
  });

  it('should throw an error for negative perPage', () => {
    const items = [1, 2, 3, 4, 5];
    const options = { perPage: -1 };
    expect(() => paginate(items, options)).toThrow('perPage must be a positive integer');
  });

  it('should throw an error for non-integer perPage', () => {
    const items = [1, 2, 3, 4, 5];
    const options = { perPage: 1.5 };
    expect(() => paginate(items, options)).toThrow('perPage must be a positive integer');
  });
});

import { ListOptions } from '../api/listOptions';

function paginate(items: any[], options: ListOptions): any[] {
  validateListOptions(options);

  if (options.perPage === 0) {
    return items;
  }

  const startIndex = (options.page - 1) * options.perPage;
  const endIndex = startIndex + options.perPage;

  return items.slice(startIndex, endIndex);
}

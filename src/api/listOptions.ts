interface ListOptions {
  page?: number;
  perPage?: number;
  sortBy?: string;
  sortOrder?: 'asc' | 'desc';
}

function validateListOptions(options: ListOptions): void {
  if (options.perPage !== undefined && (options.perPage < 0 || !Number.isInteger(options.perPage))) {
    throw new Error('perPage must be a positive integer');
  }
}

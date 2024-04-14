// utils/formatPrice.ts

/**
 * Formats a number as a price string.
 * @param price The price value to be formatted.
 * @param currency The currency symbol to use. Defaults to '$'.
 * @param locale The locale to use for formatting. Defaults to 'en-US'.
 * @returns A formatted price string.
 */
 export const formatPrice = (price: number, currency: string = '$', locale: string = 'en-US'): string => {
    return new Intl.NumberFormat(locale, { style: 'currency', currencyDisplay: 'symbol', currency }).format(price);
  };
  
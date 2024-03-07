// Define the structure for localized strings
type LocalizedStrings = {
    [locale: string]: {
      [key: string]: string;
    };
  };
  
  // Implement the localized strings
  const strings: LocalizedStrings = {
    en: {
      greeting: "Hello",
    },
    es: {
      greeting: "Hola",
    },
  };
  
  // Function to get a localized string
  export const getLocalizedString = (key: string, locale: string): string => {
    return strings[locale]?.[key] || key;
  };
  
  // Example usage
  console.log(getLocalizedString("greeting", "en")); // Output: Hello
  console.log(getLocalizedString("greeting", "es")); // Output: Hola
  
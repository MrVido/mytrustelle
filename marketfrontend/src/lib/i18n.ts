import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';
import HttpBackend from 'i18next-http-backend';
import LanguageDetector from 'i18next-browser-languagedetector';

// Define the structure of your translation files
interface ITranslation {
  authentication: {
    [key: string]: string;
  };
  commons: {
    [key: string]: any; // Adjust this to be more specific based on actual structure
  };
  app: {
    [key: string]: any; // Adjust this to be more specific based on actual structure
  };
}

// Configure i18next
i18n
  .use(HttpBackend)
  .use(LanguageDetector)
  .use(initReactI18next)
  .init({
    fallbackLng: 'en',
    debug: true,
    interpolation: {
      escapeValue: false, // Not needed for React as it escapes by default
    },
    backend: {
      loadPath: '/locales/{{lng}}/translation.json',
    },
    react: {
      useSuspense: false, // You can set to true if you want to handle loading states
    },
  });

export default i18n;
export type { ITranslation };

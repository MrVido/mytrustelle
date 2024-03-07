// utils/validators.ts

/**
 * Validates an email address format.
 * @param email The email address to validate.
 * @returns true if the email address is valid, otherwise false.
 */
 export const validateEmail = (email: string): boolean => {
    const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return regex.test(email);
  };
  
  /**
   * Checks the strength of a password.
   * @param password The password to check.
   * @returns true if the password meets the strength criteria, otherwise false.
   */
  export const isPasswordStrong = (password: string): boolean => {
    const minLength = 8;
    const regex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$/; // Example: Minimum eight characters, at least one uppercase letter, one lowercase letter, and one number
    return password.length >= minLength && regex.test(password);
  };
  
  /**
   * Validates a generic text field for not being empty.
   * @param text The text to validate.
   * @returns true if the text is not empty, otherwise false.
   */
  export const isNotEmpty = (text: string): boolean => {
    return text.trim().length > 0;
  };
  
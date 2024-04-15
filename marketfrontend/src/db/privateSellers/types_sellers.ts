import { Vehicle, PaymentHistory } from "@prisma/client";

export interface CreatePrivateSellerDto {
  userId: number; // User ID linking to the User model
  name: string;
  email: string;
  phone?: string;
  address?: string;
  city?: string;
  region?: string;
  postalCode?: string;
}

export interface UpdatePrivateSellerDto {
  name?: string;
  email?: string;
  phone?: string;
  address?: string;
  city?: string;
  region?: string;
  postalCode?: string;
}

export interface PrivateSellerQueryOptions {
  includeVehicles?: boolean;
  includePaymentHistory?: boolean;
}

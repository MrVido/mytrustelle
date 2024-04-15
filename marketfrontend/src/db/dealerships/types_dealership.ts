import { Vehicle, Subscription, Logo, OwnershipHistory, Review, PaymentHistory, ExternalIntegration, FavoriteVehicle, PriceDropNotification, VehiclePriceHistory, VehicleFinanceOption, EnvironmentalData, VehicleSafetyRating, CustomModificationRequest, CustomerFeedback, TestDriveSchedule, VehicleComparison } from '@prisma/client';

export interface CreateDealershipDto {
  groupId?: number;
  parentId?: number;
  name: string;
  countryCode: string;
  address: string;
  city: string;
  region: string;
  postalCode: string;
  descriptionEn?: string;
  descriptionFr?: string;
  phone?: string;
  fax?: string;
  url?: string;
  inventoryActive: boolean;
  active: boolean;
  slug: string;
  email?: string;
  defaultLanguage: string;
  latitude?: number;
  longitude?: number;
  seoDescription?: string;
  facebookUrl?: string;
  twitterUrl?: string;
  instagramUrl?: string;
  youtubeUrl?: string;
  linkedinUrl?: string;
  pinterestUrl?: string;
}

export interface UpdateDealershipDto {
  groupId?: number;
  parentId?: number;
  name?: string;
  countryCode?: string;
  address?: string;
  city?: string;
  region?: string;
  postalCode?: string;
  descriptionEn?: string;
  descriptionFr?: string;
  phone?: string;
  fax?: string;
  url?: string;
  inventoryActive?: boolean;
  active?: boolean;
  slug?: string;
  email?: string;
  defaultLanguage?: string;
  latitude?: number;
  longitude?: number;
  seoDescription?: string;
  facebookUrl?: string;
  twitterUrl?: string;
  instagramUrl?: string;
  youtubeUrl?: string;
  linkedinUrl?: string;
  pinterestUrl?: string;
}

export interface DealershipQueryOptions {
  includeVehicles?: boolean;
  includeSubscriptions?: boolean;
  includeLogos?: boolean;
  includeOwnershipHistory?: boolean;
  includeReviews?: boolean;
  includePaymentHistory?: boolean;
  includeExternalIntegrations?: boolean;
  includeFavoritedBy?: boolean;
  includePriceNotifications?: boolean;
  includePriceHistories?: boolean;
  includeFinanceOptions?: boolean;
  includeEnvironmentalDatas?: boolean;
  includeSafetyRatings?: boolean;
  includeCustomModRequests?: boolean;
  includeFeedbacks?: boolean;
  includeTestDrives?: boolean;
  includeVehicleComparisons?: boolean;
}

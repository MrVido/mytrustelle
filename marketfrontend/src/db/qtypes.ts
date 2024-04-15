import { OwnerType, RoleType, VehicleStatus, InterestType, NotificationStatus } from '@prisma/client';

export interface CreateRoleDto {
    type: RoleType;   // Use the RoleType enum instead of string
    permissions: RolePermissions;  // You might want to define a more specific type for permissions if needed
  }
  interface RolePermissions {
    canEdit: boolean;
    canDelete: boolean;
    canView: boolean;
    canCreate: boolean;
  }


// qtypes.ts
export interface CreateUserDto {
    username: string;
    email: string;
    name?: string;
    // Add other fields as necessary
  }
  
  export interface UpdateUserDto {
    username?: string;
    email?: string;
    name?: string;
    // Add other fields as necessary
  }
  
  

  
  export interface CreateVehicleDto {
    ownerId: number;
    ownerType: OwnerType;
    frameCategoryId: number;
    frameStyleId: number;
    vin: string;
    make: string;
    model: string;
    titleAd: string;
    trim: string;
    year: number;
    mileage: number;
    color: string;
    bodyType: string;
    fuelType: string;
    transmission: string;
    numberOfDoors: number;
    interiorFeatures?: any; // Specify more precise type if possible
    exteriorFeatures?: any; // Specify more precise type if possible
    condition: string;
    price: number;
    description?: string;
    status: VehicleStatus;
    photos?: any; // Specify more precise type if JSON or similar
    location?: string;
    performanceData?: any; // Specify more precise type if JSON or similar
    emissionDetails?: any; // Specify more precise type if JSON or similar
    warrantyInformation?: any; // Specify more precise type if JSON or similar
    listedAt?: Date;
    updatedAt?: Date;
    deleted?: boolean;
    deletedAt?: Date;
    options?: any; // Specify more precise type if JSON or similar
    securityFeatures?: any; // Specify more precise type if JSON or similar
    batteryHealth?: string;
    estimatedBatteryLife?: number;
    chargeCycles?: number;
    ecoRating?: string;
}
  
  
  export interface CreateSupplierDto {
    companyName: string;
    city: string;
    country: string;
    // Add other necessary fields
  }
  
  export interface UpdateDealershipDto {
    name?: string;
    countryCode?: string;
    // Add other fields for updates
  }
  
  export interface UpdateVehicleDto {
    make?: string;
    model?: string;
    year?: number;
    // Add other updateable fields
  }
  
  export interface CreateBookingDto {
    vehicleId: number;
    userId: number;
    scheduledDate: Date;
    status: string;
  }
  
  export interface UpdateBookingDto {
    scheduledDate?: Date;
    status?: string;
  }
  
  export interface CreateReviewDto {
    dealershipId: number;
    authorId: number;
    rating: number;
    content: string;
  }
  

  // qtypes.ts
export interface CreateDealershipDto {
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
    inventoryActive: boolean; // Assuming default values are not set in the Prisma schema
    active: boolean;         // Assuming default values are not set in the Prisma schema
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
    groupId?: number;         // Optional as per Prisma schema
    parentId?: number;        // Optional as per Prisma schema
  }
  
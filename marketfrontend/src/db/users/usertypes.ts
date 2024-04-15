import {
  RoleType as PrismaRoleType,
  PrismaClient,
  Prisma,
} from "@prisma/client";

// Enum to mirror the Prisma RoleType enum
export enum RoleType {
  SUB_USER = "SUB_USER",
  USER = "USER",
  SALES = "SALES",
  ADMIN = "ADMIN",
  DEV_USER = "DEV_USER",
  SENIOR_DEV = "SENIOR_DEV",
  CONTROLLER = "CONTROLLER",
  SUPER_ADMIN = "SUPER_ADMIN",
  KING = "KING",
}

// Permissions interface using Prisma's JsonValue for compatibility
interface RolePermissions {
  canEdit: boolean;
  canDelete: boolean;
  canView: boolean;
  canCreate: boolean;
}

// Data Transfer Object (DTO) for creating roles
interface CreateRoleDto {
  type: PrismaRoleType; // Use the RoleType from Prisma which should be aligned with your schema
  permissions: RolePermissions; // Directly use RolePermissions interface
}

// Define default permissions for each role
const defaultRolePermissions: Record<RoleType, RolePermissions> = {
  [RoleType.SUB_USER]: {
    canEdit: false,
    canDelete: false,
    canView: true,
    canCreate: false,
  },
  [RoleType.USER]: {
    canEdit: true,
    canDelete: false,
    canView: true,
    canCreate: false,
  },
  [RoleType.SALES]: {
    canEdit: true,
    canDelete: false,
    canView: true,
    canCreate: true,
  },
  [RoleType.ADMIN]: {
    canEdit: true,
    canDelete: true,
    canView: true,
    canCreate: true,
  },
  [RoleType.DEV_USER]: {
    canEdit: true,
    canDelete: false,
    canView: true,
    canCreate: true,
  },
  [RoleType.SENIOR_DEV]: {
    canEdit: true,
    canDelete: true,
    canView: true,
    canCreate: true,
  },
  [RoleType.CONTROLLER]: {
    canEdit: true,
    canDelete: true,
    canView: true,
    canCreate: true,
  },
  [RoleType.SUPER_ADMIN]: {
    canEdit: true,
    canDelete: true,
    canView: true,
    canCreate: true,
  },
  [RoleType.KING]: {
    canEdit: true,
    canDelete: true,
    canView: true,
    canCreate: true,
  },
};


import { RoleType, Prisma, PrismaClient } from '@prisma/client';

const prisma = new PrismaClient();


interface RolePermissions {
  canEdit: boolean;
  canDelete: boolean;
  canView: boolean;
  canCreate: boolean;
}

interface CreateRoleDto {
  type: RoleType;
  permissions: Prisma.JsonValue;  // Use Prisma's JsonValue for better type compatibility
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
  }
};

// Function to create a new role with permissions

export const createRole = async (roleType: RoleType) => {
    const roleData: CreateRoleDto = {
      type: roleType,
      permissions: defaultRolePermissions[roleType] as Prisma.JsonValue  // Use type assertion here
    };
  
    return await prisma.role.create({
      data: roleData,
    });
  };
// Example Usage:
const newAdminRole = createRole(RoleType.SUB_USER);
console.log(newAdminRole);

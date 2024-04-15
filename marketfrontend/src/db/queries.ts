// queries.ts
import { PrismaClient, RoleType } from '@prisma/client';


import {
  CreateUserDto,
  UpdateUserDto,
  CreateDealershipDto,
  CreateVehicleDto,
  UpdateVehicleDto,
  CreateBookingDto,
  UpdateBookingDto,
  CreateReviewDto,
  CreateRoleDto
} from './qtypes';

const prisma = new PrismaClient();

// ----------------------
// User Queries
// ----------------------

export const createUser = async (userData: CreateUserDto) => {
  return await prisma.user.create({
    data: userData,
  });
};


export const getUserById = async (id: number) => {
  return await prisma.user.findUnique({
    where: { id },
  });
};

export const updateUser = async (id: number, userData: UpdateUserDto) => {
  return await prisma.user.update({
    where: { id },
    data: userData,
  });
};

export const deleteUser = async (id: number) => {
  return await prisma.user.delete({
    where: { id },
  });
};

// ----------------------
// Role Queries
// ----------------------

export const createRole = async (roleType: RoleType) => {
    const roleData: CreateRoleDto = {
        type: roleType,
        permissions: JSON.stringify([roleType])  // Serialize to JSON string if necessary
    };

    return await prisma.role.create({
        data: roleData,
    });
};
export const getRoleById = async (id: number) => {
  return await prisma.role.findUnique({
    where: { id },
  });
};

// ----------------------
// UserRole Queries
// ----------------------

export const assignRoleToUser = async (userId: number, roleId: number) => {
  return await prisma.userRole.create({
    data: {
      userId,
      roleId,
    },
  });
};

// ----------------------
// Dealership Queries
// ----------------------

export const createDealership = async (dealershipData: CreateDealershipDto) => {
  return await prisma.dealership.create({
    data: dealershipData,
  });
};

export const getDealershipById = async (id: number) => {
  return await prisma.dealership.findUnique({
    where: { id },
  });
};

// ----------------------
// Vehicle Queries
// ----------------------

export const createVehicle = async (vehicleData: CreateVehicleDto) => {
  return await prisma.vehicle.create({
    data: vehicleData,
  });
};

export const getVehicleById = async (id: number) => {
  return await prisma.vehicle.findUnique({
    where: { id },
  });
};

export const updateVehicle = async (id: number, vehicleData: UpdateVehicleDto) => {
  return await prisma.vehicle.update({
    where: { id },
    data: vehicleData,
  });
};

export const deleteVehicle = async (id: number) => {
  return await prisma.vehicle.delete({
    where: { id },
  });
};
// ----------------------
// Booking Queries
// ----------------------

export const createBooking = async (bookingData: CreateBookingDto) => {
  return await prisma.booking.create({
    data: bookingData,
  });
};

export const updateBooking = async (id: number, bookingData: UpdateBookingDto) => {
  return await prisma.booking.update({
    where: { id },
    data: bookingData,
  });
};

export const deleteBooking = async (id: number) => {
  return await prisma.booking.delete({
    where: { id },
  });
};

// ----------------------
// Review Queries
// ----------------------

export const createReview = async (reviewData: CreateReviewDto) => {
  return await prisma.review.create({
    data: reviewData,
  });
};

// Define the exports object
const dbQueries = {
  createUser,
  getUserById,
  updateUser,
  deleteUser,
  createRole,
  getRoleById,
  assignRoleToUser,
  createDealership,
  getDealershipById,
  createVehicle,
  getVehicleById,
  updateVehicle,
  deleteVehicle
};

// Export the queries object
export default dbQueries;

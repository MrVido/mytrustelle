import { PrismaClient } from '@prisma/client';
import { CreatePrivateSellerDto, UpdatePrivateSellerDto, PrivateSellerQueryOptions } from './types_sellers';

const prisma = new PrismaClient();

// Function to create a new private seller
export async function createPrivateSeller(data: CreatePrivateSellerDto) {
    return await prisma.privateSeller.create({
        data,
    });
}

// Function to retrieve a private seller by ID
// Function to retrieve a private seller by ID with conditional inclusion of related data
export async function getPrivateSellerById(id: number, options?: PrivateSellerQueryOptions) {
    const includeOptions = {
        include: {
            vehicles: options?.includeVehicles ?? false,  // Ensure fallback to false if options is undefined
            paymentHistory: options?.includePaymentHistory ?? false  // Same here
        }
    };
    return await prisma.privateSeller.findUnique({
        where: { id },
        ...includeOptions  // Correctly apply the structured include options
    });
}


// Function to update a private seller by ID
export async function updatePrivateSeller(id: number, data: UpdatePrivateSellerDto) {
    return await prisma.privateSeller.update({
        where: { id },
        data,
    });
}

// Function to delete a private seller by ID
export async function deletePrivateSeller(id: number) {
    return await prisma.privateSeller.delete({
        where: { id },
    });
}

// Example of a more complex query, such as retrieving all private sellers in a specific city
export async function getPrivateSellersByCity(city: string) {
    return await prisma.privateSeller.findMany({
        where: {
            city,
        },
    });
}

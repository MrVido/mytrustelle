import { PrismaClient } from '@prisma/client';
import { CreatePrivateSellerDto, UpdatePrivateSellerDto } from './types_sellers';

const prisma = new PrismaClient();

// Function to create a new private seller
export async function createPrivateSeller(data: CreatePrivateSellerDto) {
    return await prisma.privateSeller.create({
        data
    });
}

// Function to update an existing private seller
export async function updatePrivateSeller(id: number, data: UpdatePrivateSellerDto) {
    return await prisma.privateSeller.update({
        where: { id },
        data
    });
}

// Function to delete a private seller
export async function deletePrivateSeller(id: number) {
    return await prisma.privateSeller.delete({
        where: { id }
    });
}

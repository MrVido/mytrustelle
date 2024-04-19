import { PrismaClient } from '@prisma/client';

const prisma = new PrismaClient();

async function main() {
  const user = await prisma.user.create({
    data: {
      username: 'user1',
      email: 'user1@example.com',
      type: 'SUB_USER', // Assuming this is correct based on your RoleType enum
      permissions: {}, // Assuming permissions is a JSON field, provide appropriate JSON object
      roles: {
        create: [{
          role: {
            connect: { name: 'SUB_USER' }, // This assumes 'SUB_USER' role is pre-seeded in the database
          },
        }],
      },
    },
  });

  console.log(`Created user with ID: ${user.id}`);
  // Add more entities similarly
}

main()
  .catch((e) => {
    console.error(e);
    process.exit(1);
  })
  .finally(async () => {
    await prisma.$disconnect();
  });

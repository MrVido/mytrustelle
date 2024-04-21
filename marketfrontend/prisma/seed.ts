import { PrismaClient } from '@prisma/client';

const prisma = new PrismaClient();

async function main() {
  // First, ensure the role exists. This is a simplistic check/creation logic.
  const roleName = 'SUB_USER';
  let role = await prisma.role.findUnique({
    where: { name: roleName },
  });

  if (!role) {
    role = await prisma.role.create({
      data: { name: roleName, permissions: {} },
    });
  }

  // Create user and connect to the existing or newly created role
  const user = await prisma.user.create({
    data: {
      username: 'user1',
      email: 'user1@example.com',
      type: roleName,
      permissions: {}, // Customize according to your needs
      roles: {
        create: [{
          role: {
            connect: { name: roleName },
          },
        }],
      },
    },
  });

  console.log(`Created user with ID: ${user.id}`);
}

main()
  .catch((e) => {
    console.error(e);
    process.exit(1);
  })
  .finally(async () => {
    await prisma.$disconnect();
  });

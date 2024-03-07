import React from 'react';
import Link from 'next/link';
import { useCategory } from '../../hooks/useCategory'; // Assuming a custom hook for category data fetching

interface Category {
  id: number;
  name: string;
}

const CategoryMenu: React.FC = () => {
  const { categories, isLoading, error } = useCategory(); // Utilize the useCategory hook

  if (isLoading) {
    return <p>Loading categories...</p>;
  }

  if (error) {
    return <p>Error fetching categories: {error.message}</p>;
  }

  return (
    <ul>
      {categories.map((category: Category) => (
        <li key={category.id}>
          <Link href={`/categories/${category.id}`}>
            {category.name}
          </Link>
        </li>
      ))}
    </ul>
  );
};

export default CategoryMenu;

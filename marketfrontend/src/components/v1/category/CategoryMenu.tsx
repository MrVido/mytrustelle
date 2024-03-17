// CategoryMenu.tsx
import React from 'react';
import Link from 'next/link';

interface Category {
  id: string;
  name: string;
  slug: string; // Used in the URL
}

interface CategoryMenuProps {
  categories: Category[];
}

const CategoryMenu: React.FC<CategoryMenuProps> = ({ categories }) => {
  return (
    <nav>
      <ul className="category-menu">
        {categories.map((category) => (
          <li key={category.id}>
            {/* Assuming the use of Next.js for navigation */}
            <Link href={`/categories/${category.slug}`}>
              <a>{category.name}</a>
            </Link>
          </li>
        ))}
      </ul>
      <style jsx>{`
        .category-menu {
          list-style: none;
          padding: 0;
          margin: 0;
          display: flex;
          flex-direction: column; // or 'row' for a horizontal menu
        }
        .category-menu li {
          margin: 0.5rem 0; // Adjust spacing as needed
        }
        .category-menu a {
          text-decoration: none;
          color: #0070f3; // Example color, adjust as needed
          transition: color 0.3s ease;
        }
        .category-menu a:hover {
          color: #0056b3; // Example hover color, adjust as needed
        }
      `}</style>
    </nav>
  );
};

export default CategoryMenu;

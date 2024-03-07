import React from 'react';

// Define the structure of a category's properties
interface CategoryProps {
  id: number;
  name: string;
  description?: string; // Optional description
  slug: string; // Useful for links or routing
}

interface CategoryListProps {
  categories: CategoryProps[]; // Array of categories
}

const CategoryList: React.FC<CategoryListProps> = ({ categories }) => {
  return (
    <div className="category-list">
      <h2>Categories</h2>
      <ul>
        {categories.map((category) => (
          <li key={category.id}>
            <a href={`/categories/${category.slug}`}>{category.name}</a>
            {category.description && <p>{category.description}</p>}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default CategoryList;

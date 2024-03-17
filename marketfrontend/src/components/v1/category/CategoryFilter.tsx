// CategoryFilter.tsx
import React from 'react';

interface Category {
  id: string;
  name: string;
}

interface CategoryFilterProps {
  categories: Category[];
  onSelect: (categoryId: string) => void;
}

const CategoryFilter: React.FC<CategoryFilterProps> = ({ categories, onSelect }) => {
  return (
    <div>
      <h3>Filter by Category</h3>
      <ul>
        {categories.map((category) => (
          <li key={category.id} onClick={() => onSelect(category.id)}>
            {category.name}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default CategoryFilter;

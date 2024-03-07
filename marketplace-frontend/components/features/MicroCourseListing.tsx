import React from 'react';

interface Course {
  id: string;
  title: string;
  description: string;
  duration: string;
  enrollmentLink: string;
}

interface MicroCourseListingProps {
  courses: Course[];
}

const MicroCourseListing: React.FC<MicroCourseListingProps> = ({ courses }) => {
  return (
    <div className="micro-course-listing">
      <h2>Micro Courses</h2>
      {courses.length > 0 ? (
        <ul>
          {courses.map((course) => (
            <li key={course.id} className="course-item">
              <h3>{course.title}</h3>
              <p>{course.description}</p>
              <p>Duration: {course.duration}</p>
              <a href={course.enrollmentLink} target="_blank" rel="noopener noreferrer">
                Enroll Now
              </a>
            </li>
          ))}
        </ul>
      ) : (
        <p>No micro-courses available at the moment. Check back later!</p>
      )}
    </div>
  );
};

export default MicroCourseListing;

import React, { useState } from 'react';
import { Camera } from 'lucide-react';

const App = () => {
  const [currentPage, setCurrentPage] = useState('home');
  const [showNewPost, setShowNewPost] = useState(false);
  const [showCategories, setShowCategories] = useState(false);
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  
  const categories = [
    'General', 'Technology', 'Sports', 'Entertainment', 'Science',
    'Health', 'Food', 'Travel', 'Fashion', 'Art', 'Music'
  ];

  const LeftSidebar = () => (
    <aside className="w-64 h-screen fixed left-0 border-r border-gray-200 bg-white">
      {!isLoggedIn ? (
        <div className="p-4 flex flex-col items-center">
          <img src="/api/placeholder/50/50" alt="logo" className="w-12 h-12 mb-4" />
          <h2 className="text-lg font-semibold mb-4">Join the conversation</h2>
          <button 
            onClick={() => setCurrentPage('login')}
            className="w-full mb-2 py-2 px-4 bg-blue-500 text-white rounded hover:bg-blue-600"
          >
            Login
          </button>
          <button 
            onClick={() => setCurrentPage('register')}
            className="w-full py-2 px-4 border border-blue-500 text-blue-500 rounded hover:bg-blue-50"
          >
            Sign up
          </button>
        </div>
      ) : (
        <div className="p-4">
          <div className="flex items-center gap-2 mb-6">
            <img src="/api/placeholder/40/40" alt="Profile" className="w-10 h-10 rounded-full" />
          </div>
          <nav className="flex flex-col gap-2">
            {[
              { id: 'home', icon: 'home', label: 'Home' },
              { id: 'categories', icon: 'filter', label: 'Categories' },
              { id: 'profile', icon: 'user', label: 'Profile' },
              { id: 'settings', icon: 'settings', label: 'Settings' }
            ].map(item => (
              <button
                key={item.id}
                onClick={() => setCurrentPage(item.id)}
                className={`flex items-center gap-2 p-2 rounded hover:bg-gray-100 ${
                  currentPage === item.id ? 'bg-blue-50 text-blue-500' : ''
                }`}
              >
                <Camera size={20} />
                <span>{item.label}</span>
              </button>
            ))}
            <button
              onClick={() => setShowNewPost(true)}
              className="mt-4 flex items-center gap-2 p-2 bg-blue-500 text-white rounded hover:bg-blue-600"
            >
              <Camera size={20} />
              <span>New Post</span>
            </button>
          </nav>
        </div>
      )}
    </aside>
  );

  const RightSidebar = () => (
    <aside className="w-64 h-screen fixed right-0 border-l border-gray-200 p-4 bg-white">
      <input
        type="text"
        placeholder="Search..."
        className="w-full p-2 border rounded mb-4"
      />
      <div className="flex gap-2 text-sm text-gray-500">
        <a href="#about" className="hover:text-blue-500">about</a>
        <span>·</span>
        <a href="#contact" className="hover:text-blue-500">contact</a>
      </div>
    </aside>
  );

  const NewPostModal = () => (
    showNewPost && (
      <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
        <div className="bg-white rounded-lg w-full max-w-lg p-4">
          <div className="flex justify-between items-center mb-4">
            <button 
              onClick={() => setShowNewPost(false)}
              className="text-gray-500 hover:text-gray-700"
            >
              Cancel
            </button>
            <button className="bg-blue-500 text-white px-4 py-1 rounded hover:bg-blue-600">
              Post
            </button>
          </div>
          <div className="flex gap-2 mb-4">
            <img src="/api/placeholder/40/40" alt="Profile" className="w-10 h-10 rounded-full" />
            <textarea
              placeholder="What's up?"
              className="w-full p-2 border rounded resize-none"
              rows={4}
            />
          </div>
          <button 
            onClick={() => setShowCategories(true)}
            className="text-blue-500 hover:text-blue-600"
          >
            Categories
          </button>
        </div>
      </div>
    )
  );

  const CategoriesModal = () => (
    showCategories && (
      <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
        <div className="bg-white rounded-lg w-full max-w-lg p-4">
          <div className="flex justify-between items-center mb-4">
            <button 
              onClick={() => setShowCategories(false)}
              className="text-gray-500 hover:text-gray-700"
            >
              Cancel
            </button>
            <h2 className="text-lg font-semibold">Choose your categories:</h2>
            <button className="bg-blue-500 text-white px-4 py-1 rounded hover:bg-blue-600">
              Done
            </button>
          </div>
          <div className="grid grid-cols-3 gap-2">
            {categories.map(category => (
              <button
                key={category}
                className="p-2 border rounded hover:bg-blue-50 hover:border-blue-500"
              >
                {category}
              </button>
            ))}
          </div>
        </div>
      </div>
    )
  );

  const MainContent = () => (
    <main className="ml-64 mr-64 min-h-screen p-4">
      <div className="flex justify-center mb-4">
        <img src="/api/placeholder/50/50" alt="logo" className="w-12 h-12" />
      </div>
      {/* Content will change based on currentPage */}
      <div className="max-w-2xl mx-auto">
        {currentPage === 'home' && <div>Home Content</div>}
        {currentPage === 'categories' && <div>Categories Content</div>}
        {currentPage === 'profile' && <div>Profile Content</div>}
        {currentPage === 'settings' && <div>Settings Content</div>}
        {currentPage === 'login' && <div>Login Form</div>}
        {currentPage === 'register' && <div>Register Form</div>}
      </div>
    </main>
  );

  return (
    <div className="h-screen bg-white">
      <LeftSidebar />
      <MainContent />
      <RightSidebar />
      <NewPostModal />
      <CategoriesModal />
    </div>
  );
};

export default App;
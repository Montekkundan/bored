// File: @/types/user.ts

// Represents the response from authentication endpoints (login/register)
export interface AuthResponse {
    token: string;
    user: User;
  }
  
  // Represents a user in the system
  export interface User {
    id: number;
    username: string;
    email: string;
    bio: string;
    interests: string[] | null;
    latitude: number;
    longitude: number;
    profile_picture: string;
    cover_photo: string;
    social_links: SocialLinks;
    oauth_providers: string[] | null;
    audio_enabled: boolean;
    video_enabled: boolean;
    roles: UserRole[];
    email_verified: boolean;
    phone_number: string;
    phone_verified: boolean;
    two_factor_enabled: boolean;
    reward_points: number;
    created_at: string;
    updated_at: string;
  }
  
  // Represents the social links of a user
  export interface SocialLinks {
    facebook?: string;
    twitter?: string;
    instagram?: string;
    linkedin?: string;
    github?: string;
    [key: string]: string | undefined;
  }
  
  // Enum for user roles
  export enum UserRole {
    USER = 'user',
    ADMIN = 'admin',
    MODERATOR = 'moderator'
  }
  
  // Represents the data needed to create a new user
  export interface CreateUserDto {
    username: string;
    email: string;
    password: string;
  }
  
  // Represents the data needed to update a user
  export interface UpdateUserDto {
    username?: string;
    email?: string;
    bio?: string;
    interests?: string[];
    latitude?: number;
    longitude?: number;
    profile_picture?: string;
    cover_photo?: string;
    social_links?: SocialLinks;
    audio_enabled?: boolean;
    video_enabled?: boolean;
    phone_number?: string;
  }
  
  // Represents the data needed for user login
  export interface LoginDto {
    email: string;
    password: string;
  }
  
  // Represents the data needed for user registration
  export interface RegisterDto {
    username: string;
    email: string;
    password: string;
  }
  
  // Represents a paginated response of users
  export interface PaginatedUsersResponse {
    data: User[];
    total: number;
    page: number;
    limit: number;
  }
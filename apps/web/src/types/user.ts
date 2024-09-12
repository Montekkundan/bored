export interface AuthResponse {
  token: string;
  user: User;
}

export interface UserResponse {
  data: User;
  status: string;
}

export interface User {
  id: number;
  username: string;
  email: string;
  token_version: number;
  bio: string;
  interests: string[] | null;
  latitude: number;
  longitude: number;
  profile_picture: string;
  cover_photo: string;
  social_links: string | SocialLinks;
  oauth_providers: string[] | null;
  audio_enabled: boolean;
  video_enabled: boolean;
  roles: UserRole[];
  email_verified: boolean;
  phone_number: string;
  phone_verified: boolean;
  two_factor_enabled: boolean;
  reward_points: number;
  Followers: any | null; // Replace 'any' with a more specific type if available
  Following: any | null; // Replace 'any' with a more specific type if available
  notifications: any | null; // Replace 'any' with a more specific type if available
  created_at: string;
  updated_at: string;
  chats: any | null; // Replace 'any' with a more specific type if available
  moderation_votes: any | null; // Replace 'any' with a more specific type if available
}

export interface SocialLinks {
  facebook?: string;
  twitter?: string;
  instagram?: string;
  linkedin?: string;
  github?: string;
  [key: string]: string | undefined;
}

export enum UserRole {
  USER = 'user',
  ADMIN = 'admin',
  MODERATOR = 'moderator',
}

export interface CreateUserDto {
  username: string;
  email: string;
  password: string;
}

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

export interface LoginDto {
  email: string;
  password: string;
}

export interface RegisterDto {
  username: string;
  email: string;
  password: string;
}

export interface PaginatedUsersResponse {
  data: User[];
  total: number;
  page: number;
  limit: number;
}

export interface LoginResponse {
  data: {
    token: {
      access_token: string;
      refresh_token: string;
    };
    user: User;
  };
  message: string;
  status: string;
}

export const ALL_USER_ROLES = [
  UserRole.USER,
  UserRole.ADMIN,
  UserRole.MODERATOR,
];

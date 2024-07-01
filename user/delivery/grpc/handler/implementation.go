package handler

import (
	"context"

	"github.com/TheManhattanProject/crypt_pg_backend/user/delivery/grpc/pb"
	userdomain "github.com/TheManhattanProject/crypt_pg_backend/user/domain/user"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) HealthCheck(
	ctx context.Context,
	in *pb.HealthCheckReq,
) (*pb.HealthCheckRes, error) {
	return &pb.HealthCheckRes{
		Ok: true,
	}, nil
}

func (s *server) SignUp(ctx context.Context, in *pb.SignupReq) (*pb.SignupRes, error) {
	var err error

	user := &userdomain.User{}
	if in.CountryCode != nil {
		user.CountryCode = *in.CountryCode
		user.PhoneNumber = in.InputField
	} else {
		user.Email = in.InputField
	}

	user, err = s.usecases.User.CreateOrUpdateUser(ctx, *user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	// Send the OTP
	err = s.usecases.User.SendOTP(ctx, user.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.SignupRes{
		UserId: user.ID.String(),
	}, nil
}

func (s *server) VerifyUser(ctx context.Context, in *pb.VerifyUserReq) (*pb.VerifyUserRes, error) {
	userId, err := uuid.Parse(in.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid uuid %v", in.UserId)
	}

	isVerified, err := s.usecases.User.ValidateOTP(ctx, userId, string(in.Code))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.VerifyUserRes{
		Verified: isVerified,
	}, nil
}

func (s *server) Onboard(ctx context.Context, in *pb.OnboardingReq) (*pb.OnboardingRes, error) {
	userId, err := uuid.Parse(in.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid uuid %v", in.UserId)
	}

	user := &userdomain.User{
		ID:       userId,
		UserName: in.UserName,
		Password: in.Password,
	}

	user, err = s.usecases.User.CreateOrUpdateUser(ctx, *user)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &pb.OnboardingRes{
		Ok: true,
	}, nil
}

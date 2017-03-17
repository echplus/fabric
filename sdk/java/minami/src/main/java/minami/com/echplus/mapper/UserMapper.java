package minami.com.echplus.mapper;

import java.util.List;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.annotations.Select;

import minami.com.echplus.domain.UserDomain;

@Mapper
public interface UserMapper {

    @Select("SELECT row, id, enrollmentId FROM Users WHERE row <= #{row}")
    List<UserDomain> getUsers0(@Param("row") int row);

    List<UserDomain> getUsers1(Integer row);
    
    @Insert("INSERT INTO Users (row, id, enrollmentId) VALUES (#{row}, #{id}, #{enrollmentId})")
    int newUser(UserDomain user);

}
